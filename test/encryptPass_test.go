package test

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func Test_encryptPass(t *testing.T) {
	password := "hahha"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return

	}

	fmt.Println(&hashedPassword)
}
