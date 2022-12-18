package config

import (
	"database/sql"
	"fmt"
	"github.com/MCPutro/Go-MyWallet/keys"
	_ "github.com/lib/pq"
	"log"
	"time"
)

//var ENV = ".env.gcp"

func InitDatabase() (*sql.DB, error) {
	////run in localhost
	//err := godotenv.Load(keys.Environment)
	//
	//mustGetEnv := func(k string) string {
	//	v := os.Getenv(k)
	//	if v == "" {
	//		log.Fatalf("Warning: %s environment variable not set.", k)
	//	}
	//	return v
	//}
	//
	//var (
	//	dbUser = mustGetEnv("DB_USER")     //postgres
	//	dbPass = mustGetEnv("DB_PASSWORD") //welcome1
	//	dbHost = mustGetEnv("DB_HOSTNAME") //localhost
	//	dbName = mustGetEnv("DB_NAME")     //postgres
	//	dbPort = mustGetEnv("DB_PORT")     //5432
	//	dbSSL  = mustGetEnv("DB_SSL")      //disable or require
	//)

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		keys.DbHost, keys.DbUser, keys.DbPass, keys.DbPort, keys.DbName, keys.DbSSL)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("error create db connection, err: %s", err)
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	//DB_Selected = "postgres"

	return db, nil
}
