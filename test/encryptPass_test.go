package test

import (
	"fmt"
	"github.com/MCPutro/Go-MyWallet/app"
	"github.com/MCPutro/Go-MyWallet/helper"
	"github.com/MCPutro/Go-MyWallet/service"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
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
	jwtService := service.NewJwtService("aaa")

	token := jwtService.GenerateToken("12345", "")

	fmt.Println(token)
}

func TestEncryption(t *testing.T) {
	data := "testencriptdata"

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	app.EncryptionDecryptionKey = os.Getenv("ENCRYPTION_DECRYPTION_KEY")

	encryption := helper.Encryption(data)

	fmt.Println(encryption)
}

func TestDecryption(t *testing.T) {
	data := "vrLRo359vmzQ1q7m8qNhIUClleL6RYAC25opcGPB4q3d+R1Xrxt8YsNvug=="

	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	app.EncryptionDecryptionKey = os.Getenv("ENCRYPTION_DECRYPTION_KEY")

	decryption := helper.Decryption(data)

	fmt.Println(decryption)
}
