package test

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/service"
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

func TestJWT(t *testing.T) {
	jwtService := service.NewJwtService("aaa", "bbb")

	token := jwtService.GenerateToken("12345")

	fmt.Println(token)
}
