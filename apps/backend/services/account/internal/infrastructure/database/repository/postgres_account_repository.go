package repository

import (
	"context"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
	"github.com/StephenFeddes/memorytrivia/account/internal/infrastructure/database"
)

type PostgresAccountRepository struct {
	db *database.PostgresDB
}

func NewPostgresAccountRepository(db *database.PostgresDB) *PostgresAccountRepository {
	return &PostgresAccountRepository{db: db}
}

// CreateAccount creates a new account in the database and returns the account ID.
// It returns an error if the account could not be created.
func (r *PostgresAccountRepository) CreateAccount(
	ctx context.Context, username string, email string, password string,
) (uint32, error) {
	var accountID uint32
	err := r.db.Client.QueryRow("INSERT INTO account (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		username, email, password).Scan(&accountID)

	if err != nil {
		return 0, err
	}
	return accountID, nil
}

// GetAccountByID retrieves an account from the database by its email.
// It returns an error if the account could not be found.
func (r *PostgresAccountRepository) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Client.QueryRow(
		"SELECT id, username, email, password FROM account WHERE email = $1",
		email).Scan(&account.ID, &account.Username, &account.Email, &account.Password)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountByUsername retrieves an account from the database by its username.
// It returns an error if the account could not be found.
func (r *PostgresAccountRepository) GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.Client.QueryRow(
		"SELECT id, username, email, password FROM account WHERE username = $1",
		username).Scan(&account.ID, &account.Username, &account.Email, &account.Password)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
