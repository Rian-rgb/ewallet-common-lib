package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func SetupLogger(env string) {
	log = logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      env == "development",
	})

	log.SetOutput(os.Stdout)

	if env == "production" {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Info("Logger successfully initiated")
}

func Info(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Error(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

// WithContext digunakan untuk microservices distributed tracing.
// Mengambil requestId dari context value.
func WithContext(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		return log.WithFields(logrus.Fields{})
	}

	if reqID, ok := ctx.Value("request_id").(string); ok {
		return log.WithField("request_id", reqID)
	}

	return log.WithFields(logrus.Fields{})
}
