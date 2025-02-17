package util

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes input password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword compares input password with hashed password
func ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func IsValidPassword(password string) bool {
	//contraseña de al menos 8 caracteres

	if len(password) < 8 {
		return false
	}

	return true
}
