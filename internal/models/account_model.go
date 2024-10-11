package models

import "time"

type Account struct {
	ID            int     `db:"id" json:"id"`
	AccountNumber string  `db:"account_number" json:"account_number"`
	AccountType   string  `db:"account_type" json:"account_type"`
	Balance       float64 `db:"balance" json:"balance"`
	Frozen        bool    `db:"frozen" json:"frozen"`
}

type AccountType string

const (
	Checking AccountType = "checking"
	Savings  AccountType = "savings"
)

type Account struct {
	ID          int         `json:"id" db:"id"`
	AccountType AccountType `json:"account_type" db:"account_type"`
	Balance     float64     `json:"balance" db:"balance"`
	Status      string      `json:"status" db:"status"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}
