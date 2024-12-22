package main

import (
	"context"
	"github.com/go-micro/plugins/v4/logger/logrus"
	lg "github.com/sirupsen/logrus"
	"go-micro.dev/v4/logger"
)

// CustomLogger wraps the standard logger.
type CustomLogger struct {
	traceId string
	spanId  string
	log logger.Logger
}

// NewCustomLogger creates a new instance of CustomLogger.
func NewCustomLogger() *CustomLogger {
	// Initialize a logrus logger
	log := logrus.WithJSONFormatter(&lg.JSONFormatter{})

	// Wrap the logrus logger with the go-micro logger
	return &CustomLogger{
		log: logrus.NewLogger(log), // Create a go-micro logger using logrus
	}
}

// WithCTX creates a new CustomLogger with the provided context.
func (cl *CustomLogger) WithCTX(ctx context.Context) *CustomLogger {
	// In this case, we don't need to store context directly inside the logger.
	if ctx != nil {
		// Just an example: retrieve a value from the context (like request ID)
		if traceId, ok := ctx.Value("trace_id").(string); ok {
			// Log context value with the log message
			cl.traceId = traceId
		}
		if spanId, ok := ctx.Value("span_id").(string); ok {
			// Log context value with the log message
			cl.spanId = spanId
		}
	}
	return cl
}
func (cl *CustomLogger) Info(args ...interface{}) {
	cl.log.Fields(map[string]interface{}{"trace_id": cl.traceId, "span_id": cl.spanId}).Log(logger.InfoLevel, args...)
}

func (cl *CustomLogger) Error(args ...interface{}) {
	cl.log.Fields(map[string]interface{}{"trace_id": cl.traceId, "span_id": cl.spanId}).Log(logger.ErrorLevel, args...)
}

func (cl *CustomLogger) Infof(format string, args ...interface{}) {
	cl.log.Fields(map[string]interface{}{"trace_id": cl.traceId, "span_id": cl.spanId}).Logf(logger.InfoLevel, format, args...)
}

func (cl *CustomLogger) Errorf(format string, args ...interface{}) {
	cl.log.Fields(map[string]interface{}{"trace_id": cl.traceId, "span_id": cl.spanId}).Logf(logger.ErrorLevel, format, args...)
}

func main() {
	// Example of creating a loggercd
	//customLogger := NewCustomLogger()
	customLogger := NewCustomLogger()

	// Example context with a value (like a request ID)
	ctx := context.WithValue(context.Background(), "trace_id", "12345")
	ctx = context.WithValue(ctx, "span_id", "AABBCC")

	// Log using the custom logger with context
	customLogger.WithCTX(ctx).Info("test Info with context")

	// Use WithCTX to create a new CustomLogger (in this case, WithCTX is just a no-op here)
	customLogger.WithCTX(ctx).Error("test Error with new context")

}
