package log

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger will be used by other packages to print log entries.
var logger *logrus.Logger

func init() {
	logger = logrus.New()
	if os.Getenv("FIAAS_ENVIRONMENT") != "" {
		logger.SetReportCaller(true)
		logger.Out = os.Stdout
		logger.Formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "caller",
			},
		}
	}
}

func Logger() *logrus.Logger {
	return logger
}

// SetLevel sets the log level for the logger.
func SetLevel(levelStr string) error {
	lvl, err := logrus.ParseLevel(levelStr)
	if err != nil {
		return err
	}
	logger.Level = lvl
	return nil
}

// Level returns the currently set level of the logger.
func Level() string {
	return logger.Level.String()
}
