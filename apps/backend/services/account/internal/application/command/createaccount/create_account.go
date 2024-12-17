package createaccount

import (
	"context"
	"errors"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
)

// Interfaces for the dependencies of the CreateAccount use case.
type databaseAccountCreator interface {
	CreateAccount(ctx context.Context, username string, email string, password string) (uint32, error)
	GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error)
	GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error)
}
type bCryptManager interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hash string) error
}
type jwtManager interface {
	GenerateToken(id uint32, username string, email string) (string, error)
	ValidateToken(token string) (*entity.Account, error)
}

// CreateAccount represents the use case for creating a new account.
type CreateAccount struct {
	db           databaseAccountCreator // The database to create the account in.
	bCryptHasher bCryptManager          // The BCrypt hasher to hash the password.
	jwtManager   jwtManager             // The JWT manager to generate and validate tokens.
}

// NewCreateAccount returns a new CreateAccount with the given database, bCrypt, and JWT dependencies.
func NewCreateAccount(db databaseAccountCreator, b bCryptManager, j jwtManager) *CreateAccount {
	return &CreateAccount{
		db:           db,
		bCryptHasher: b,
		jwtManager:   j,
	}
}

// Execute creates a new account with the given username, email, and password and returns a jwt token.
// It returns an error if the account could not be created due to invalid input or username/email already existing.
func (c *CreateAccount) Execute(ctx context.Context, username string, email string, password string) (string, uint32, error) {
	// Check if the email or username already exists.
	_, err := c.db.GetAccountByUsername(ctx, username)
	if err == nil {
		return "", 0, errors.New("username already exists")
	}
	_, err = c.db.GetAccountByEmail(ctx, email)
	if err == nil {
		return "", 0, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := c.bCryptHasher.HashPassword(password)
	if err != nil {
		return "", 0, errors.New("password could not be hashed")
	}

	// Create the account in the database
	accountID, err := c.db.CreateAccount(ctx, username, email, hashedPassword)
	if err != nil {
		return "", accountID, errors.New("database error")
	}

	// Generate a JWT token for the account and return it
	token, err := c.jwtManager.GenerateToken(accountID, username, email)
	if err != nil {
		return "", accountID, errors.New("jwt error")
	}

	return token, accountID, nil
}
