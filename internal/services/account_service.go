package services

import (
	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"github.com/kudaibergenoff/mybankapi/internal/repositories"
)

type AccountService struct {
	repo repositories.AccountRepositoryInterface
}

func NewAccountService(repo repositories.AccountRepositoryInterface) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.repo.Create(account)
}

func (s *AccountService) UpdateAccount(account *models.Account) error {
	return s.repo.Update(account)
}

func (s *AccountService) DeleteAccount(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *AccountService) GetAccountByID(id uuid.UUID) (*models.Account, error) {
	return s.repo.FindByID(id)
}

func (s *AccountService) FreezeAccount(id uuid.UUID) error {
	return s.repo.FreezeAccount(id)
}

func (s *AccountService) UnfreezeAccount(id uuid.UUID) error {
	return s.repo.UnfreezeAccount(id)
}
