package repository

import (
	"database/sql"
	"time"

	"github.com/gabrielalmir/go-gateway-api/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (r *AccountRepository) Save(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO account(id, name, email, api_key, balance, created_at, updated_at)
		VALUES($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		account.ID,
		account.Name,
		account.Email,
		account.ApiKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByApiKey(apiKey string) (*domain.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM account WHERE api_key = $1
	`)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(apiKey)

	account := &domain.Account{}
	err = row.Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) FindById(id string) (*domain.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, name, email, api_key, balance, created_at, updated_at
		FROM account WHERE id = $1
	`)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	account := &domain.Account{}
	err = row.Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.ApiKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	var currentBalance float64
	err = tx.QueryRow(`SELECT balance FROM account WHERE id = $1 FOR UPDATE`, account.ID).
		Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		UPDATE accounts
		SET balance = $1, updated_at = $2
		WHERE id = $3
	`, account.Balance, time.Now(), account.ID)

	if err != nil {
		return err
	}

	return tx.Commit()
}
