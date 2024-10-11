package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kudaibergenoff/mybankapi/internal/models"
	"time"
)

type AccountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(account *models.Account) error {
	account.ID = uuid.New()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	_, err := r.db.NamedExec(`INSERT INTO accounts (id, account_type, balance, is_frozen, created_at, updated_at) VALUES (:id, :account_type, :balance, :is_frozen, :created_at, :updated_at)`, account)
	return err
}

func (r *AccountRepository) Update(account *models.Account) error {
	account.UpdatedAt = time.Now()
	_, err := r.db.NamedExec(`UPDATE accounts SET account_type=:account_type, balance=:balance, is_frozen=:is_frozen, updated_at=:updated_at WHERE id=:id`, account)
	return err
}

func (r *AccountRepository) FreezeAccount(id uuid.UUID) error {
	_, err := r.db.Exec(`UPDATE accounts SET is_frozen = true, updated_at = $1 WHERE id = $2`, time.Now(), id)
	return err
}

func (r *AccountRepository) UnfreezeAccount(id uuid.UUID) error {
	_, err := r.db.Exec(`UPDATE accounts SET is_frozen = false, updated_at = $1 WHERE id = $2`, time.Now(), id)
	return err
}

func (r *AccountRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec(`DELETE FROM accounts WHERE id = $1`, id)
	return err
}

func (r *AccountRepository) FindByID(id uuid.UUID) (*models.Account, error) {
	var account models.Account
	err := r.db.Get(&account, "SELECT * FROM accounts WHERE id = $1", id)
	return &account, err
}
