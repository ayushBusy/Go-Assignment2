package services

import (
	"errors"
	"fmt"

	"banking_system/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) Create(account *models.Account) error {
	return s.db.Create(account).Error
}

func (s *AccountService) GetByID(id uint) (*models.Account, error) {
	var account models.Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *AccountService) GetAccountDetail(id uint) (*models.AccountDetail, error) {
	var account models.Account
	if err := s.db.First(&account, id).Error; err != nil {
		return nil, err
	}

	var accountCustomers []models.AccountCustomer
	if err := s.db.Preload("Customer").Where("account_id = ?", id).Find(&accountCustomers).Error; err != nil {
		return nil, err
	}

	detail := &models.AccountDetail{
		ID:            account.ID,
		AccountNumber: account.AccountNumber,
		BranchID:      account.BranchID,
		AccountType:   account.AccountType,
		Interest:      account.Interest,
		Balance:       account.Balance,
		CreatedAt:     account.CreatedAt,
		Customers:     make([]models.CustomerInfo, 0),
	}

	for _, ac := range accountCustomers {
		detail.Customers = append(detail.Customers, models.CustomerInfo{
			CustomerID: ac.CustomerID,
			FirstName:  ac.Customer.FirstName,
			LastName:   ac.Customer.LastName,
			Email:      ac.Customer.Email,
			Phone:      ac.Customer.Phone,
			Role:       ac.Role,
		})
	}

	return detail, nil
}

func (s *AccountService) GetAll() ([]models.Account, error) {
	var accounts []models.Account
	if err := s.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountService) Update(account *models.Account) error {
	return s.db.Save(account).Error
}

func (s *AccountService) Delete(id uint) error {
	return s.db.Delete(&models.Account{}, id).Error
}

func (s *AccountService) AddCustomer(accountID, customerID uint) (*models.AccountDetail, error) {
	var account models.Account
	if err := s.db.First(&account, accountID).Error; err != nil {
		return nil, fmt.Errorf("account not found: %w", err)
	}

	var customer models.Customer
	if err := s.db.First(&customer, customerID).Error; err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	//this checks if customer is already linked to this account
	var existingLink models.AccountCustomer
	if err := s.db.Where("account_id = ? AND customer_id = ?", accountID, customerID).First(&existingLink).Error; err == nil {
		return nil, errors.New("customer is already linked to this account")
	}

	var count int64
	if err := s.db.Model(&models.AccountCustomer{}).Where("account_id = ?", accountID).Count(&count).Error; err != nil {
		return nil, fmt.Errorf("failed to count existing customers: %w", err)
	}

	//this handles the role based on customer count
	role := "primary_holder"
	if count > 0 {
		role = "joint_holder"
		// updates account type to 'joint' when adding second customer
		if err := s.db.Model(&account).Update("account_type", "joint").Error; err != nil {
			return nil, fmt.Errorf("failed to update account type: %w", err)
		}
	}

	link := models.AccountCustomer{
		AccountID:  accountID,
		CustomerID: customerID,
		Role:       role,
	}

	if err := s.db.Create(&link).Error; err != nil {
		return nil, fmt.Errorf("failed to add customer: %w", err)
	}

	return s.GetAccountDetail(accountID)
}

func (s *AccountService) RemoveCustomer(accountID, customerID uint) error {
	var linkCount int64
	if err := s.db.Model(&models.AccountCustomer{}).Where("account_id = ?", accountID).Count(&linkCount).Error; err != nil {
		return fmt.Errorf("failed to count customers: %w", err)
	}

	if err := s.db.Delete(&models.AccountCustomer{}, "account_id = ? AND customer_id = ?", accountID, customerID).Error; err != nil {
		return err
	}

	if linkCount == 2 {
		if err := s.db.Model(&models.Account{}).Where("id = ?", accountID).Update("account_type", "savings").Error; err != nil {
			return fmt.Errorf("failed to update account type: %w", err)
		}
	}
	return nil
}

func (s *AccountService) GetTransactions(accountID uint) ([]models.Transaction, error) {
	var txs []models.Transaction
	if err := s.db.Where("account_id = ?", accountID).Order("created_at asc").Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}

func (s *AccountService) Deposit(accountID uint, amount float64, description string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	var txRecord *models.Transaction

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var account models.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, accountID).Error; err != nil {
			return err
		}

		account.Balance += amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		newTx := models.Transaction{
			AccountID:   accountID,
			Type:        "deposit",
			Amount:      amount,
			Description: description,
		}
		if err := tx.Create(&newTx).Error; err != nil {
			return err
		}
		txRecord = &newTx
		return nil
	})

	if err != nil {
		return nil, err
	}

	return txRecord, nil
}

func (s *AccountService) Withdraw(accountID uint, amount float64, description string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	var txRecord *models.Transaction

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var account models.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, accountID).Error; err != nil {
			return err
		}

		if account.Balance < amount {
			return errors.New("insufficient balance")
		}

		account.Balance -= amount
		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		newTx := models.Transaction{
			AccountID:   accountID,
			Type:        "withdrawal",
			Amount:      amount,
			Description: description,
		}
		if err := tx.Create(&newTx).Error; err != nil {
			return err
		}
		txRecord = &newTx
		return nil
	})

	if err != nil {
		return nil, err
	}

	return txRecord, nil
}

