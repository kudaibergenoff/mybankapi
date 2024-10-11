package services

import (
	_ "errors"
	"testing"

	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	_ "github.com/kudaibergenoff/mybankapi/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDebitAccount(t *testing.T) {
	mockAccountRepo := new(MockAccountRepository)
	mockTransactionRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockAccountRepo, mockTransactionRepo)

	accountID := uuid.New()
	account := &models.Account{
		ID:       accountID,
		IsFrozen: false,
		Balance:  1000.00,
	}
	mockAccountRepo.On("FindByID", accountID).Return(account, nil)
	mockTransactionRepo.On("Create", mock.Anything).Return(nil)
	mockTransactionRepo.On("UpdateBalance", accountID, -100.00).Return(nil)

	err := service.DebitAccount(accountID, 100.00)

	assert.NoError(t, err)
	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestCreditAccount(t *testing.T) {
	mockAccountRepo := new(MockAccountRepository)
	mockTransactionRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockAccountRepo, mockTransactionRepo)

	accountID := uuid.New()
	account := &models.Account{
		ID:       accountID,
		IsFrozen: false,
		Balance:  1000.00,
	}
	mockAccountRepo.On("FindByID", accountID).Return(account, nil)
	mockTransactionRepo.On("Create", mock.Anything).Return(nil)
	mockTransactionRepo.On("UpdateBalance", accountID, 100.00).Return(nil)

	err := service.CreditAccount(accountID, 100.00)

	assert.NoError(t, err)
	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestTransferFunds(t *testing.T) {
	mockAccountRepo := new(MockAccountRepository)
	mockTransactionRepo := new(MockTransactionRepository)
	service := NewTransactionService(mockAccountRepo, mockTransactionRepo)

	fromAccountID := uuid.New()
	toAccountID := uuid.New()
	fromAccount := &models.Account{
		ID:       fromAccountID,
		IsFrozen: false,
		Balance:  1000.00,
	}
	toAccount := &models.Account{
		ID:       toAccountID,
		IsFrozen: false,
		Balance:  500.00,
	}

	mockAccountRepo.On("FindByID", fromAccountID).Return(fromAccount, nil)
	mockAccountRepo.On("FindByID", toAccountID).Return(toAccount, nil)
	mockTransactionRepo.On("Create", mock.Anything).Return(nil)
	mockTransactionRepo.On("UpdateBalance", fromAccountID, -100.00).Return(nil)
	mockTransactionRepo.On("UpdateBalance", toAccountID, 100.00).Return(nil)

	err := service.TransferFunds(fromAccountID, toAccountID, 100.00)

	assert.NoError(t, err)
	mockAccountRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}
