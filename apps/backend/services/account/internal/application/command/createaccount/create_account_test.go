package createaccount

import (
	"context"
	"errors"
	"testing"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
)

func TestCreateAccount(t *testing.T) {
	t.Run("should create an account", func(t *testing.T) {
		// Arrange
		mockDB := &mockDatabase{
			GetAccountByEmailError:    errors.New("account not found"),
			GetAccountByUsernameError: errors.New("account not found"),
		}
		createAccount := NewCreateAccount(mockDB, &mockBcryptManager{}, &mockJWTManager{})
		username := "testUser1"
		email := "testuser@example.com"
		hashedPassword := "hashedpassword"

		// Act
		jwtToken, _, err := createAccount.Execute(context.TODO(), username, email, hashedPassword)

		// Assert
		if mockDB.LastArgs.Username != username {
			t.Errorf("Expected username %v, got %v", username, mockDB.LastArgs.Username)
		}
		if mockDB.LastArgs.Email != email {
			t.Errorf("Expected email %v, got %v", email, mockDB.LastArgs.Email)
		}
		if mockDB.LastArgs.Password != hashedPassword {
			t.Errorf("Expected hashed password %v, got %v", hashedPassword, mockDB.LastArgs.Password)
		}
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if jwtToken != "jwttoken" {
			t.Errorf("Expected jwttoken %v, got %v", "account created", jwtToken)
		}
	})

	tests := []struct {
		name          string
		mockDB        *mockDatabase
		mockBCrypt    *mockBcryptManager
		mockJWT       *mockJWTManager
		username      string
		email         string
		password      string
		expectedError string
	}{
		{
			name: "should fail to create an account if email is taken",
			mockDB: &mockDatabase{
				GetAccountByUsernameError: errors.New("account not found"),
			},
			mockBCrypt:    &mockBcryptManager{},
			mockJWT:       &mockJWTManager{},
			username:      "testUser1",
			email:         "testuser@example.com",
			password:      "secureP@ssword1",
			expectedError: "email already exists",
		},
		{
			name: "should fail to create an account if username is taken",
			mockDB: &mockDatabase{
				GetAccountByEmailError: errors.New("account not found"),
			},
			mockBCrypt:    &mockBcryptManager{},
			mockJWT:       &mockJWTManager{},
			username:      "testUser1",
			email:         "testuser@example.com",
			password:      "secureP@ssword1",
			expectedError: "username already exists",
		},
		{
			name: "should fail to create an account if password can't be hashed",
			mockDB: &mockDatabase{
				GetAccountByEmailError:    errors.New("account not found"),
				GetAccountByUsernameError: errors.New("account not found"),
			},
			mockBCrypt:    &mockBcryptManager{LastBCryptHashError: errors.New("password could not be hashed")},
			mockJWT:       &mockJWTManager{},
			username:      "testUser1",
			email:         "testuser@example.com",
			password:      "secureP@ssword1",
			expectedError: "password could not be hashed",
		},
		{
			name: "should fail to create an account if there is a database error",
			mockDB: &mockDatabase{
				GetAccountByEmailError:    errors.New("account not found"),
				GetAccountByUsernameError: errors.New("account not found"),
				DatabaseError:             errors.New("database error"),
			},
			mockBCrypt:    &mockBcryptManager{},
			mockJWT:       &mockJWTManager{},
			username:      "testUser1",
			email:         "testuser@example.com",
			password:      "secureP@ssword1",
			expectedError: "database error",
		},
		{
			name: "should fail to create an account if the jwt token cannot be created",
			mockDB: &mockDatabase{
				GetAccountByEmailError:    errors.New("account not found"),
				GetAccountByUsernameError: errors.New("account not found"),
			},
			mockBCrypt:    &mockBcryptManager{},
			mockJWT:       &mockJWTManager{LastJWTError: errors.New("jwt error")},
			username:      "testUser1",
			email:         "testuser@example.com",
			password:      "secureP@ssword1",
			expectedError: "jwt error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			createAccount := NewCreateAccount(tt.mockDB, tt.mockBCrypt, tt.mockJWT)

			// Act
			_, _, err := createAccount.Execute(context.TODO(), tt.username, tt.email, tt.password)

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
	LastJWTError error
}

func (j *mockJWTManager) GenerateToken(id uint32, username string, email string) (string, error) {
	return "jwttoken", j.LastJWTError
}
func (j *mockJWTManager) ValidateToken(token string) (*entity.Account, error) {
	return &entity.Account{
		ID:       1,
		Username: "testuser",
		Email:    "testuser@example.com",
	}, nil
}

type mockBcryptManager struct {
	LastBCryptHashError error
}

func (bc *mockBcryptManager) HashPassword(password string) (string, error) {
	return "hashedpassword", bc.LastBCryptHashError
}
func (bc *mockBcryptManager) VerifyPassword(password, hash string) error {
	return nil
}

type mockDatabase struct {
	LastArgs struct {
		Username string
		Email    string
		Password string
	}
	GetAccountByEmailError    error
	GetAccountByUsernameError error
	DatabaseError             error
}

func (db *mockDatabase) CreateAccount(ctx context.Context, username string, email string, password string) (uint32, error) {
	db.LastArgs.Username = username
	db.LastArgs.Email = email
	db.LastArgs.Password = password
	return 1, db.DatabaseError
}
func (db *mockDatabase) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	return nil, db.GetAccountByEmailError
}
func (db *mockDatabase) GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error) {
	return nil, db.GetAccountByUsernameError
}
