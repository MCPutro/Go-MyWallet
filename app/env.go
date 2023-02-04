package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	/* Database */
	DbUser = mustGetEnv("DB_USER")     //postgres
	DbPass = mustGetEnv("DB_PASSWORD") //password
	DbHost = mustGetEnv("DB_HOSTNAME") //localhost
	DbName = mustGetEnv("DB_NAME")     //postgres
	DbPort = mustGetEnv("DB_PORT")     //5432
	DbSSL  = mustGetEnv("DB_SSL")      //disable or require

	/* secret key */
	EncryptionDecryptionKey = mustGetEnv("ENCRYPTION_DECRYPTION_KEY")
	JwtSecretKey            = mustGetEnv("JWT_SECRET_KEY")

	/* port for application */
	AppPort = mustGetEnv("PORT")

	WEB         = mustGetEnv("WEB_KEY")
	MOB_ANDROID = mustGetEnv("MOB_ANDROID_KEY")
)

func mustGetEnv(k string) string {
	/* if run in localhost will load from variable(file) environment */
	err := godotenv.Load(".env")
	if err != nil {
		return ""
	}

	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.", k)
	}
	return v
}
