package loginaccount

import (
	"context"
	"errors"
	"testing"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
)

func TestLoginAccount(t *testing.T) {
	t.Run("should login an account and return a jwt", func(t *testing.T) {
		// Arrange
		mockDB := &mockDatabase{}
		loginAccount := NewLoginAccount(mockDB, &mockBcryptManager{}, &mockJWTManager{})
		username := "testUser"
		email := "testUser@gmail.com"
		password := "hashedPassw0rd"

		// Act
		jwt, _, err := loginAccount.Execute(context.TODO(), username, email, password)

		// Assert
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if jwt != "jwttoken" {
			t.Errorf("Expected jwttoken, got %v", jwt)
		}
	})

	tests := []struct {
		name          string
		mockDB        *mockDatabase
		mockBCrypt    *mockBcryptManager
		mockJWT       *mockJWTManager
		expectedError string
	}{
		{
			name: "should fail to login if user does not exist",
			mockDB: &mockDatabase{
				GetAccountByUsernameError: errors.New("account not found"),
				GetAccountByEmailError:    errors.New("account not found"),
			},
			mockBCrypt:    &mockBcryptManager{},
			mockJWT:       &mockJWTManager{},
			expectedError: "username/email or password is incorrect",
		},
		{
			name:   "should fail to login if password is incorrect",
			mockDB: &mockDatabase{},
			mockBCrypt: &mockBcryptManager{
				ValidationError: errors.New("passwords do not match"),
			},
			mockJWT:       &mockJWTManager{},
			expectedError: "username/email or password is incorrect",
		},
		{
			name:       "should fail to login if the jwt token could not be generated",
			mockDB:     &mockDatabase{},
			mockBCrypt: &mockBcryptManager{},
			mockJWT: &mockJWTManager{
				jwtError: errors.New("jwt error"),
			},
			expectedError: "jwt error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			loginAccount := NewLoginAccount(tt.mockDB, tt.mockBCrypt, tt.mockJWT)
			username := "testUser"
			email := "testUser@gmail.com"
			password := "hashedPassw0rd"

			// Act
			_, _, err := loginAccount.Execute(context.TODO(), username, email, password)

			// Assert
			if err == nil {
				t.Fatalf("Expected an error, got nil")
			}
			if err.Error() != tt.expectedError {
				t.Errorf("Expected error '%s', got '%v'", tt.expectedError, err)
			}
		})
	}
}

type mockJWTManager struct {
	jwtError error
}

func (j *mockJWTManager) GenerateToken(id uint32, username string, email string) (string, error) {
	return "jwttoken", j.jwtError
}
func (j *mockJWTManager) ValidateToken(token string) (*entity.Account, error) {
	return &entity.Account{
		ID:       1,
		Username: "testUser",
		Email:    "testUser@gmail.com",
		Password: "hashedpassword",
	}, nil
}

type mockBcryptManager struct {
	ValidationError error
}

func (bc *mockBcryptManager) HashPassword(password string) (string, error) {
	return "hashedpassword", nil
}
func (bc *mockBcryptManager) VerifyPassword(password, hash string) error {
	return bc.ValidationError
}

type mockDatabase struct {
	GetAccountByEmailError    error
	GetAccountByUsernameError error
}

func (db *mockDatabase) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	return &entity.Account{
		ID:       1,
		Username: "testuser",
		Email:    "testUser@gmail.com",
		Password: "hashedPassw0rd",
	}, db.GetAccountByEmailError
}
func (db *mockDatabase) GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error) {
	return &entity.Account{
		ID:       1,
		Username: "testuser",
		Email:    "testUser@gmail.com",
		Password: "hashedPassw0rd",
	}, db.GetAccountByUsernameError
}
