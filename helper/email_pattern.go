package helper

import "regexp"

func IsEmail(email string) bool {
	//_, err := mail.ParseAddress(email)
	compile := regexp.MustCompile(`^[a-zA-Z0-9._\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`)

	return compile.MatchString(email)
	//return err == nil
}
