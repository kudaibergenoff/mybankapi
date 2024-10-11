package repositories

import "github.com/kudaibergenoff/mybankapi/internal/models"

type AccountRepository interface {
}

type accountRepository struct {
	models.Account
}

func NewAccountRepository() AccountRepository {
	return accountRepository{}
}

func (a *accountRepository) Find(id string) (*models.Account, error) {
	return nil, nil
}
