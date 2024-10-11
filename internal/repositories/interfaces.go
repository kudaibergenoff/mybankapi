package repositories

import (
	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
)

type AccountRepositoryInterface interface {
	Create(account *models.Account) error
	Update(account *models.Account) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.Account, error)
	FreezeAccount(id uuid.UUID) error
	UnfreezeAccount(id uuid.UUID) error
}

type TransactionRepositoryInterface interface {
	Create(transaction *models.Transaction) error
	UpdateBalance(accountID uuid.UUID, amount float64) error
	FindAccountByID(id uuid.UUID) (*models.Account, error)
}
