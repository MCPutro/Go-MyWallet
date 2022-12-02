package helper

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnClose(conn *sql.Conn) {
	conn.Close()
	fmt.Println("---->>> conn close")

}

func CommitOrRollback(err error, tx *sql.Tx) {
	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			log.Println("[ERROR] failed Rollback : ", err2)
		} else {
			log.Println("[INFO] success Rollback")
		}
	} else {
		if err2 := tx.Commit(); err2 != nil {
			log.Println("[ERROR] failed Commit : ", err2)
		} else {
			log.Println("[INFO] success Commit")
		}
	}
}
