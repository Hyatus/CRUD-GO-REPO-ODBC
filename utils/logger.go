package utils

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger(){
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{}) // Set the log format to JSON

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fmt.Println("Logger initialized, logging to app.log")
	fmt.Print(err)
	if err == nil {
		Log.SetOutput(file)
	}else{
		Log.Info("Failed to log to file, using default stderr")
		Log.SetOutput(os.Stderr) // Fallback to stderr if file cannot be opened
	}
	Log.SetLevel(logrus.InfoLevel) // Set the log level to Info
}