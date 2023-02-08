package zapLog

import (
	"context"

	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"go.uber.org/zap"
)

func (z zapLogger) Printf(format string, args ...interface{}) {
	z.log.Infof(format, args...)
}

func (z zapLogger) Info(ctx context.Context, message string, attrs ...logger.Attribute) {
	z.setAttributes(attrs...).Info(message)
}

func (z zapLogger) Debug(ctx context.Context, message string, attrs ...logger.Attribute) {
	z.setAttributes(attrs...).Debug(message)
}

func (z zapLogger) Warn(ctx context.Context, message string, attrs ...logger.Attribute) {
	z.setAttributes(attrs...).Warn(message)
}

func (z zapLogger) Error(ctx context.Context, message string, attrs ...logger.Attribute) {
	z.setAttributes(attrs...).Error(message)
}

func (z zapLogger) setAttributes(attrs ...logger.Attribute) *zap.SugaredLogger {
	usedLog := z.log

	for _, attr := range attrs {
		usedLog = usedLog.With(attr.Key(), attr.Value())
	}

	return usedLog
}
