package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var StdLogger *log.Logger

func InitLogger() {
	logPath, err := filepath.Abs("./log/api.log")
	if err != nil {
		fmt.Println("Error reading path: ", err)
	}

	generalLog, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	multiWriter := io.MultiWriter(generalLog, os.Stdout)
	StdLogger = log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
}
