package util

import "regexp"

func IsValidEmail(email string) bool {
	//validamos que sea un email correcto

	format := "^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9]+$"

	return regexp.MustCompile(format).MatchString(email)

}
