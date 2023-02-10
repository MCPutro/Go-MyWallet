package app

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func InitLog(path string) *logrus.Logger {
	logger := logrus.New()

	logger.SetReportCaller(true)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("[ERROR] Initial Logrus : ", err)
		return nil
	}
	logger.SetOutput(file)

	return logger
}
