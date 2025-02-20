package util

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type UtilService struct {
}

func NewUtilService() *UtilService {
	return &UtilService{}
}

// HashPassword hashes input password using bcrypt
func (u *UtilService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword compares input password with hashed password
func (u *UtilService) ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UtilService) IsValidPassword(password string) bool {
	//contraseña de al menos 8 caracteres

	if len(password) < 8 {
		return false
	}

	return true
}

func (u *UtilService) IsValidEmail(email string) bool {
	//validamos que sea un email correcto

	format := "^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9]+$"

	return regexp.MustCompile(format).MatchString(email)

}
