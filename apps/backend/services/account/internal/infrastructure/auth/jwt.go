package auth

import (
	"fmt"
	"time"

	"github.com/StephenFeddes/memorytrivia/account/internal/domain/entity"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret: secret}
}

// GenerateToken creates a token with a specified expiration time
func (j *JWT) GenerateToken(id uint32, username string, email string) (string, error) {
	const duration = time.Hour * 24
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Add claims
	claims["id"] = id
	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(duration).Unix() // Set expiration time

	return token.SignedString([]byte(j.secret))
}

// ValidateToken validates the token and checks expiration
func (j *JWT) ValidateToken(tokenString string) (*entity.Account, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	// Check if parsing failed
	if err != nil {
		return nil, err
	}

	// Check token validity
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract claims
		id := uint32(claims["id"].(float64))
		username := claims["username"].(string)
		email := claims["email"].(string)
		return &entity.Account{ID: id, Username: username, Email: email}, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}