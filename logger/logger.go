package logger

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	Log = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}
