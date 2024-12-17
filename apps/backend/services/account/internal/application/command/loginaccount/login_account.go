package loginaccount

import (
	"context"
	"errors"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
)

type databaseAccountGetter interface {
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

type LoginAccount struct {
	db            databaseAccountGetter
	bCryptManager bCryptManager
	jwtManager jwtManager
}

func NewLoginAccount(db databaseAccountGetter, b bCryptManager, j jwtManager) *LoginAccount {
	return &LoginAccount{db: db, bCryptManager: b, jwtManager: j}
}

func (l *LoginAccount) Execute(ctx context.Context, username string, email string, password string) (string, uint32, error) {
	const errMsg = "username/email or password is incorrect"

	// Get the account by username or email
	account, err := l.db.GetAccountByUsername(ctx, username)
	if err != nil {
		account, err = l.db.GetAccountByEmail(ctx, email)
		if err != nil {
			return "", 0, errors.New(errMsg)
		}
	}

	// Verify the password
	err = l.bCryptManager.VerifyPassword(password, account.Password)
	if err != nil {
		return "", account.ID, errors.New(errMsg)
	}

	// Generate a JWT token for the account and return it
	token, err := l.jwtManager.GenerateToken(account.ID, account.Username, account.Email)
	if err != nil {
		return "", account.ID, errors.New("jwt error")
	}

	return token, account.ID, nil
}
