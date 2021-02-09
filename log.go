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
				"@version":           "1", // apparently this is required on fiaas logs
			},
		}
	}
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

// Debug proxies logrus.Logger.Debug
func Debug(args ...interface{}) { logger.Debug(args...) }

// Debugf proxies logrus.Logger.Debugf
func Debugf(format string, args ...interface{}) { logger.Debugf(format, args...) }

// Error proxies logrus.Logger.Error
func Error(args ...interface{}) { logger.Error(args...) }

// Errorf proxies logrus.Logger.Errorf
func Errorf(format string, args ...interface{}) { logger.Errorf(format, args...) }

// Fatal proxies logrus.Logger.Fatal
func Fatal(args ...interface{}) { logger.Fatal(args...) }

// Fatalf proxies logrus.Logger.Fatalf
func Fatalf(format string, args ...interface{}) { logger.Fatalf(format, args...) }

// Info proxies logrus.Logger.Info
func Info(args ...interface{}) { logger.Info(args...) }

// Infof proxies logrus.Logger.Infof
func Infof(format string, args ...interface{}) { logger.Infof(format, args...) }

// Trace proxies logrus.Logger.Trace
func Trace(args ...interface{}) { logger.Trace(args...) }

// Tracef proxies logrus.Logger.Tracef
func Tracef(format string, args ...interface{}) { logger.Tracef(format, args...) }

// Warn proxies logrus.Logger.Warn
func Warn(args ...interface{}) { logger.Warn(args...) }

// Warnf proxies logrus.Logger.Warnf
func Warnf(format string, args ...interface{}) { logger.Warnf(format, args...) }
