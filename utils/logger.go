package utils

import (
	"log"
	"os"
)

var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func LogError(err error) {
	Error.Println(err)
}

func LogInfo(msg string) {
	Info.Println(msg)
}
