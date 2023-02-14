package app

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
)

func InitLog(path string) *logrus.Logger {

	//check path already exists or not
	if _, err := os.Stat(path); os.IsNotExist(err) {
		temp := strings.Split(path, "/")

		test := strings.Join(temp[:len(temp)-1], "/")
		err = os.MkdirAll(test, 0666)
		if err != nil {
			fmt.Println("error init log")
			return nil
		}
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("[ERROR] Initial Logrus : ", err)
		return nil
	}

	logger := logrus.New()

	logger.SetReportCaller(true)

	logger.SetOutput(file)

	return logger
}
