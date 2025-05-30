package logger

import (
	"log"
	"os"
)

var (
	Logger  *log.Logger
	LogFile *os.File
)

func Init_log() {
	LogFile, err := os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(LogFile, "prefix: ", log.LstdFlags)

}
