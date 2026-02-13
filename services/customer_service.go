package services

import (
	"banking_system/models"

	"gorm.io/gorm"
)

type CustomerService struct {
	db *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{db: db}
}

func (s *CustomerService) Create(customer *models.Customer) error {
	return s.db.Create(customer).Error
}

func (s *CustomerService) GetByID(id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := s.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (s *CustomerService) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	if err := s.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *CustomerService) Update(customer *models.Customer) error {
	return s.db.Save(customer).Error
}

func (s *CustomerService) Delete(id uint) error {
	return s.db.Delete(&models.Customer{}, id).Error
}

func (s *CustomerService) GetAccounts(customerID uint) ([]models.Account, error) {
	var accounts []models.Account
	err := s.db.Table("accounts").
		Joins("JOIN account_customers ON account_customers.account_id = accounts.id").
		Where("account_customers.customer_id = ?", customerID).
		Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *CustomerService) GetLoans(customerID uint) ([]models.Loan, error) {
	var loans []models.Loan
	if err := s.db.Where("customer_id = ?", customerID).Find(&loans).Error; err != nil {
		return nil, err
	}
	return loans, nil
}

