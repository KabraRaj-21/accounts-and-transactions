package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"
	"transaction/internal/logger"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Log request details
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore body

		logger.WithContext(c.Request.Context()).
			Infof("Incoming Request: %s %s Headers: %v Body: %s\n", c.Request.Method, c.Request.URL.String(), c.Request.Header, string(bodyBytes))

		// Capture response
		responseBody := &bytes.Buffer{}
		writer := &bodyWriter{ResponseWriter: c.Writer, body: responseBody}
		c.Writer = writer

		// Process request
		c.Next()

		// Log response details
		duration := time.Since(startTime)
		logger.WithContext(c.Request.Context()).
			Infof("Response: %d %s Duration: %vms Body: %s", c.Writer.Status(), http.StatusText(c.Writer.Status()), duration.Milliseconds(), responseBody.String())
	}
}

// Custom response writer to capture response body
type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture response body
	return w.ResponseWriter.Write(b)
}
