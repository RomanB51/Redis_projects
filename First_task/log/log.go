package logger

import (
	"log"
	"os"
)

var (
	Log     *log.Logger
	logFile *os.File
)

func Init_log() {
	logFile, err := os.OpenFile("log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Log = log.New(logFile, "prefix: ", log.LstdFlags)

}

func Print(str string) {
	Log.Print(str)
}

func PrintErr(str string, err error) {
	Log.Print(str, err)
}

func GetLogFile() *os.File {
	return logFile
}
