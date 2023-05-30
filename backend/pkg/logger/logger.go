package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	green   string = "\033[32m"
	yellow  string = "\033[33m"
	red     string = "\033[31m"
	reset   string = "\033[0m"
	magenta string = "\033[35m"
)

type LogLevel struct {
	infoLog    *log.Logger
	ErrorLog   *log.Logger
	warningLog *log.Logger
	debugLog   *log.Logger
}

const (
	info    = "INFO"
	errLog  = "ERROR"
	warning = "WARNING"
	debug   = "DEBUG"
)

func New() *LogLevel {
	logFile := getLogFileName()
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, file)
	return &LogLevel{
		infoLog:    log.New(multiWriter, fmt.Sprintf("%s[INFO]   \t%s", green, reset), log.Ldate|log.Ltime),
		ErrorLog:   log.New(multiWriter, fmt.Sprintf("%s[ERROR]  \t%s", red, reset), log.Ldate|log.Ltime|log.Lshortfile),
		warningLog: log.New(multiWriter, fmt.Sprintf("%s[WARNING]\t%s", yellow, reset), log.Ldate|log.Ltime),
		debugLog:   log.New(multiWriter, fmt.Sprintf("%s[DEBUG]  \t%s", magenta, reset), log.Ldate|log.Ltime),
	}
}

func (l *LogLevel) Info(message string) {
	l.infoLog.Println(message)
}
func (l *LogLevel) Error(message string) {
	l.ErrorLog.Println(message)
}
func (l *LogLevel) Warning(message string) {
	l.warningLog.Println(message)
}
func (l *LogLevel) Debug(message string) {
	l.debugLog.Println(message)
}

func getLogFileName() string {
	sessionID := getSessionID()
	dir := "./pkg/logger/logs"
	file := fmt.Sprintf("logfile_%s.txt", sessionID)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Println("error when creating dir", err)
		}
	}

	path := filepath.Join(dir, file)
	return path
}

func getSessionID() string {
	currentTime := time.Now().Format("2006-01-02_15:04")
	return currentTime
}
