package services

import (
	"banking_system/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{db: db}
}

func (s *TransactionService) Create(txn *models.Transaction) error {
	return s.db.Create(txn).Error
}

func (s *TransactionService) GetByID(id uint) (*models.Transaction, error) {
	var txn models.Transaction
	if err := s.db.First(&txn, id).Error; err != nil {
		return nil, err
	}
	return &txn, nil
}

func (s *TransactionService) GetAll() ([]models.Transaction, error) {
	var txns []models.Transaction
	if err := s.db.Find(&txns).Error; err != nil {
		return nil, err
	}
	return txns, nil
}

func (s *TransactionService) Update(txn *models.Transaction) error {
	return s.db.Save(txn).Error
}

func (s *TransactionService) Delete(id uint) error {
	return s.db.Delete(&models.Transaction{}, id).Error
}

func (s *TransactionService) GetAccountTransactions(accountID uint) ([]models.Transaction, error) {
	var txs []models.Transaction
	if err := s.db.Where("account_id = ?", accountID).Order("created_at asc").Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}

func (s *TransactionService) ProcessTransaction(accountID uint, transactionType string, amount float64, description string) (*models.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	if transactionType == "" {
		return nil, errors.New("transaction_type is required (deposit, withdrawal, transfer, etc.)")
	}

	var txRecord *models.Transaction

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var account models.Account
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&account, accountID).Error; err != nil {
			return err
		}

		// Handle withdrawal and ensure sufficient balance
		switch transactionType {
		case "withdrawal", "transfer":
			if account.Balance < amount {
				return errors.New("insufficient balance")
			}
			account.Balance -= amount
		case "deposit":
			account.Balance += amount
		default:
			return errors.New("invalid transaction_type: must be deposit, withdrawal, or transfer")
		}

		if err := tx.Save(&account).Error; err != nil {
			return err
		}

		newTx := models.Transaction{
			AccountID:   accountID,
			Type:        transactionType,
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

func (s *TransactionService) Deposit(accountID uint, amount float64, description string) (*models.Transaction, error) {
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

func (s *TransactionService) Withdraw(accountID uint, amount float64, description string) (*models.Transaction, error) {
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
