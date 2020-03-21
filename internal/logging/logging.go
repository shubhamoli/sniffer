package logging


import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	FATAL = "Fatal"
	INFO = "Info"
	WARN = "Warn"
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
	Logger.Formatter = &logrus.TextFormatter{ForceColors: false, FullTimestamp: true}
}

