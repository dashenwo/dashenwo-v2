package logger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

type Option func(*Options)

type Options struct {
	// The logging level the logger should log at. default is `InfoLevel`
	Level Level
	// fields to always be logged
	Fields map[string]interface{}
	// It's common to set this to a file, or leave it default which is `os.Stderr`
	Out io.Writer
	// Caller skip frame count for file:line info
	CallerSkipCount int
	// Alternative options
	Context context.Context
}

// WithFields set default fields for the logger
func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

// WithLevel set default level for the logger
func WithLevel(level Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

// WithOutput set default output writer for the logger
func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

// WithCallerSkipCount set frame count to skip
func WithCallerSkipCount(c int) Option {
	return func(args *Options) {
		args.CallerSkipCount = c
	}
}

func SetOption(k, v interface{}) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}

type callerSkipKey struct{}

func WithCallerSkip(i int) Option {
	return SetOption(callerSkipKey{}, i)
}

type configKey struct{}

// WithConfig pass zap.Config to logger
func WithConfig(c zap.Config) Option {
	return SetOption(configKey{}, c)
}

type encoderConfigKey struct{}

// WithEncoderConfig pass zapcore.EncoderConfig to logger
func WithEncoderConfig(c zapcore.EncoderConfig) Option {
	return SetOption(encoderConfigKey{}, c)
}

type namespaceKey struct{}

func WithNamespace(namespace string) Option {
	return SetOption(namespaceKey{}, namespace)
}