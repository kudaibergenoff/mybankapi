package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"time"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	transaction.CreatedAt = time.Now()
	_, err := r.db.NamedExec(`INSERT INTO transactions (account_id, type, amount, created_at) VALUES (:account_id, :type, :amount, :created_at)`, transaction)
	return err
}

func (r *TransactionRepository) UpdateBalance(accountID uuid.UUID, amount float64) error {
	_, err := r.db.Exec(`UPDATE accounts SET balance = balance + $1, updated_at = $2 WHERE id = $3`, amount, time.Now(), accountID)
	return err
}

func (r *TransactionRepository) FindAccountByID(id uuid.UUID) (*models.Account, error) {
	var account models.Account
	err := r.db.Get(&account, "SELECT * FROM accounts WHERE id = $1", id)
	return &account, err
}
