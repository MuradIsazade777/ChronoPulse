package logger

import (
    "log"
    "os"
)

var (
    infoLog  *log.Logger
    errorLog *log.Logger
    debugLog *log.Logger
)

func Init() {
    infoLog = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
    errorLog = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
    debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
}

func Info(msg string) {
    infoLog.Println(msg)
}

func Error(msg string) {
    errorLog.Println(msg)
}

func Debug(msg string) {
    debugLog.Println(msg)
}
