package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/app"
	"io"
)

//func getKey(k string) string {
//	//if run in localhost
//	err := godotenv.Load(env.Environment)
//	if err != nil {
//		return "nil"
//	}
//
//	v := os.Getenv(k)
//	if v == "" {
//		log.Fatalf("Warning: %s environment variable not set.", k)
//	}
//	return v
//}

func Encryption(t string) string {
	//c, err := aes.NewCipher([]byte(getKey("ENCRYPTION_DECRYPTION_KEY")))
	c, err := aes.NewCipher([]byte(app.EncryptionDecryptionKey))
	if err != nil {
		fmt.Println("1", err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println("2", err)
		return ""
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("3", err)
		return ""
	}

	b := gcm.Seal(nonce, nonce, []byte(t), nil)
	//fmt.Println(b)
	//fmt.Println(string(b))
	//fmt.Println(base64.StdEncoding.EncodeToString(b))

	return base64.StdEncoding.EncodeToString(b)
}

func Decryption(t string) string {

	tByte, err := base64.StdEncoding.DecodeString(t)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	//c, err := aes.NewCipher([]byte(getKey("ENCRYPTION_DECRYPTION_KEY")))
	c, err := aes.NewCipher([]byte(app.EncryptionDecryptionKey))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	nonceSize := gcm.NonceSize()
	if len(tByte) < nonceSize {
		fmt.Println(err)
		return ""
	}

	nonce, ciphertext := tByte[:nonceSize], tByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//fmt.Println(">>", string(plaintext))

	return string(plaintext)
}
