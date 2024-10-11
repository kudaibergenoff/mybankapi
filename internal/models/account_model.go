package models

import (
	"github.com/google/uuid"
	"time"
)

type AccountType string

const (
	Checking AccountType = "checking"
	Savings  AccountType = "savings"
)

type Account struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	AccountType AccountType `json:"account_type" db:"account_type"`
	Balance     float64     `json:"balance" db:"balance"`
	IsFrozen    bool        `json:"is_frozen" db:"is_frozen"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}
