package auth

import "golang.org/x/crypto/bcrypt"

type BCrypt struct {}
func NewBCrypt() *BCrypt {
	return &BCrypt{}
}
func (b *BCrypt) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
func (b *BCrypt) VerifyPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}