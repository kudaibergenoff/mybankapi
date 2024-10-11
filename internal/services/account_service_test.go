package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	account := &models.Account{
		AccountType: "checking",
		Balance:     1000.00,
	}
	mockRepo.On("Create", account).Return(nil)

	err := service.CreateAccount(account)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	account := &models.Account{
		ID:          uuid.New(),
		AccountType: "savings",
		Balance:     1500.00,
	}
	mockRepo.On("Update", account).Return(nil)

	err := service.UpdateAccount(account)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	mockRepo.On("Delete", id).Return(nil)

	err := service.DeleteAccount(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAccountByID(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	account := &models.Account{
		ID:          id,
		AccountType: "checking",
		Balance:     1000.00,
	}
	mockRepo.On("FindByID", id).Return(account, nil)

	result, err := service.GetAccountByID(id)

	assert.NoError(t, err)
	assert.Equal(t, account, result)
	mockRepo.AssertExpectations(t)
}

func TestFreezeAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	mockRepo.On("FreezeAccount", id).Return(nil)

	err := service.FreezeAccount(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUnfreezeAccount(t *testing.T) {
	mockRepo := new(MockAccountRepository)
	service := NewAccountService(mockRepo)

	id := uuid.New()
	mockRepo.On("UnfreezeAccount", id).Return(nil)

	err := service.UnfreezeAccount(id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
