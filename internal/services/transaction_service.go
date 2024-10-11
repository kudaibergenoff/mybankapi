package services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"github.com/kudaibergenoff/mybankapi/internal/repositories"
)

type TransactionService struct {
	accountRepo     repositories.AccountRepositoryInterface
	transactionRepo repositories.TransactionRepositoryInterface
}

func NewTransactionService(accountRepo repositories.AccountRepositoryInterface, transactionRepo repositories.TransactionRepositoryInterface) *TransactionService {
	return &TransactionService{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

func (s *TransactionService) DebitAccount(accountID uuid.UUID, amount float64) error {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return err
	}

	if account.IsFrozen {
		return errors.New("account is frozen")
	}

	if account.Balance < amount {
		return errors.New("insufficient funds")
	}

	transaction := &models.Transaction{
		AccountID: accountID,
		Type:      "debit",
		Amount:    amount,
	}

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return err
	}

	return s.transactionRepo.UpdateBalance(accountID, -amount)
}

func (s *TransactionService) CreditAccount(accountID uuid.UUID, amount float64) error {
	account, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return err
	}

	if account.IsFrozen {
		return errors.New("account is frozen")
	}

	transaction := &models.Transaction{
		AccountID: accountID,
		Type:      "credit",
		Amount:    amount,
	}

	err = s.transactionRepo.Create(transaction)
	if err != nil {
		return err
	}

	return s.transactionRepo.UpdateBalance(accountID, amount)
}

func (s *TransactionService) TransferFunds(fromAccountID, toAccountID uuid.UUID, amount float64) error {
	fromAccount, err := s.accountRepo.FindByID(fromAccountID)
	if err != nil {
		return err
	}

	if fromAccount.IsFrozen {
		return errors.New("from account is frozen")
	}

	toAccount, err := s.accountRepo.FindByID(toAccountID)
	if err != nil {
		return err
	}

	if toAccount.IsFrozen {
		return errors.New("to account is frozen")
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient funds")
	}

	debitTransaction := &models.Transaction{
		AccountID: fromAccountID,
		Type:      "debit",
		Amount:    amount,
	}

	creditTransaction := &models.Transaction{
		AccountID: toAccountID,
		Type:      "credit",
		Amount:    amount,
	}

	err = s.transactionRepo.Create(debitTransaction)
	if err != nil {
		return err
	}

	err = s.transactionRepo.UpdateBalance(fromAccountID, -amount)
	if err != nil {
		return err
	}

	err = s.transactionRepo.Create(creditTransaction)
	if err != nil {
		return err
	}

	return s.transactionRepo.UpdateBalance(toAccountID, amount)
}
