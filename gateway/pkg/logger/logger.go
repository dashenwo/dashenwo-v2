// Package log provides a log interface
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	// Default logger
	DefaultLogger Logger = NewDefaultLogger()
)

func NewDefaultLogger() Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Encoding = "console"
	productionEncoderConfig := zap.NewProductionEncoderConfig()
	productionEncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		nanos := t.Format("2006-01-02 15:04:05")
		encoder.AppendByteString([]byte(nanos))
	}
	l, _ := NewLogger(
		WithCallerSkip(2),
		WithConfig(zapConfig),
		WithEncoderConfig(productionEncoderConfig),
	)
	return l
}

// Logger is a generic logging interface
type Logger interface {
	// Init initialises options
	Init(options ...Option) error
	// The Logger options
	Options() Options
	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) Logger
	// Log writes a log entry
	Log(level Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level Level, format string, v ...interface{})
	// String returns the name of logger
	String() string
}

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) Logger {
	return DefaultLogger.Fields(fields)
}

func Log(level Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}
