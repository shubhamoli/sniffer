package logging

import (
	"os"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.SetOutput(os.Stdout)

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		// Set Info level as a default
		logLevel = logrus.InfoLevel
	}

	Logger.SetLevel(logLevel)
	Logger.Formatter = &logrus.TextFormatter{ForceColors: true, FullTimestamp: true}
}


