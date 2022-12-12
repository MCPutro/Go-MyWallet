package keys

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	environment = ".env"

	DbUser                  = mustGetEnv("DB_USER")     //postgres
	DbPass                  = mustGetEnv("DB_PASSWORD") //password
	DbHost                  = mustGetEnv("DB_HOSTNAME") //localhost
	DbName                  = mustGetEnv("DB_NAME")     //postgres
	DbPort                  = mustGetEnv("DB_PORT")     //5432
	DbSSL                   = mustGetEnv("DB_SSL")      //disable or require
	EncryptionDecryptionKey = mustGetEnv("ENCRYPTION_DECRYPTION_KEY")
	JwtSecretKey            = mustGetEnv("JWT_SECRET_KEY")

	WEB         = "cca422e771e711edb7f200ffe49cc121"
	MOB_ANDROID = "491a5e2c57394ed9b233dbf233a3174f"
)

func mustGetEnv(k string) string {
	//run in localhost will load from variable(file) environment
	err := godotenv.Load(environment)
	if err != nil {
		return ""
	}

	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.", k)
	}
	return v
}
