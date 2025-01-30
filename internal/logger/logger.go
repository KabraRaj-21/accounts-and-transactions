package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var defaultLogger = logrus.New()

func init() {
	defaultLogger.SetOutput(os.Stdout)
	defaultLogger.SetFormatter(&logrus.JSONFormatter{})
	defaultLogger.SetLevel(logrus.DebugLevel)
}

func WithContext(ctx context.Context) *logrus.Entry {
	requestID, _ := ctx.Value("request_id").(string)
	return defaultLogger.WithFields(logrus.Fields{
		"request_id": requestID,
	})
}
