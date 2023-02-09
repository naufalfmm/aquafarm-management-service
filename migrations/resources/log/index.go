package log

import (
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"github.com/naufalfmm/aquafarm-management-service/utils/logger/zapLog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLog() (logger.Logger, error) {
	return zapLog.NewZap(
		zapLog.WithTimeEncoder(zapcore.RFC3339TimeEncoder),
		zapLog.WithZapOption(zap.AddStacktrace(zap.ErrorLevel)),
	)
}
