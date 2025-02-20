package ports

type UtilService interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hashedPassword string) error
	IsValidPassword(password string) bool
	IsValidEmail(email string) bool
}
