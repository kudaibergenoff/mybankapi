package models

import (
	"github.com/google/uuid"
	"time"
)

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

type Transaction struct {
	ID        uuid.UUID       `json:"id" db:"id"`
	AccountID uuid.UUID       `json:"account_id" db:"account_id"`
	Type      TransactionType `json:"type" db:"type"`
	Amount    float64         `json:"amount" db:"amount"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}
