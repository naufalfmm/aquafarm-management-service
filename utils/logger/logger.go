package logger

import "context"

type Logger interface {
	Printf(format string, args ...interface{})

	Info(ctx context.Context, message string, attrs ...Attribute)
	Debug(ctx context.Context, message string, attrs ...Attribute)
	Warn(ctx context.Context, message string, attrs ...Attribute)
	Error(ctx context.Context, message string, attrs ...Attribute)
}
