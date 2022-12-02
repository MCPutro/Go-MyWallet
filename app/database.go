package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func InitDatabase() (*sql.DB, error) {
	dbUser := "postgres"
	dbPass := "welcome1"
	dbHost := "localhost"
	dbName := "postgres"
	dbPort := "5432"
	dbSSL := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbPort, dbName, dbSSL)

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
