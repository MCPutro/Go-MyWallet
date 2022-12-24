package helper

import (
	"database/sql"
	"log"
)

func ConnClose(conn *sql.Conn) {
	err := conn.Close()
	if err != nil {
		log.Println("[ERROR] failed close connection : ", err)
	} else {
		log.Println("[INFO] success close connection")
	}
}

func CommitOrRollback(err error, tx *sql.Tx) {
	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			log.Println("[ERROR] failed rollback : ", err2)
		} else {
			log.Println("[INFO] success rollback")
		}
	} else {
		if err2 := tx.Commit(); err2 != nil {
			log.Println("[ERROR] failed commit : ", err2)
		} else {
			log.Println("[INFO] success commit")
		}
	}
}

func Close(err error, tx *sql.Tx, con *sql.Conn) {
	CommitOrRollback(err, tx)
	ConnClose(con)
}
