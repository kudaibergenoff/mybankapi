package services

import (
	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) FindByID(id uuid.UUID) (*models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *MockAccountRepository) Create(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) Update(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAccountRepository) FreezeAccount(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockAccountRepository) UnfreezeAccount(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Create(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockTransactionRepository) UpdateBalance(accountID uuid.UUID, amount float64) error {
	args := m.Called(accountID, amount)
	return args.Error(0)
}

func (m *MockTransactionRepository) FindAccountByID(id uuid.UUID) (*models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Account), args.Error(1)
}
