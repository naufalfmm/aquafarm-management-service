package zapLog

import (
	"github.com/naufalfmm/aquafarm-management-service/utils/logger"
	"go.uber.org/zap"
)

type (
	zapLogger struct {
		log     *zap.SugaredLogger
		options options
	}
)

func NewZap(opts ...Option) (logger.Logger, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt.Apply(&o)
	}

	zapConf := zap.NewDevelopmentConfig()
	zapConf.EncoderConfig.TimeKey = "timestamp"
	zapConf.EncoderConfig.CallerKey = "caller"
	zapConf.EncoderConfig.LevelKey = "level"
	zapConf.EncoderConfig.MessageKey = "message"
	zapConf.EncoderConfig.StacktraceKey = "stacktrace"
	zapConf.Encoding = "json"
	zapConf.EncoderConfig.EncodeTime = o.timeEncoder

	log, err := zapConf.Build(o.zapOptions...)
	if err != nil {
		return nil, err
	}

	return &zapLogger{
		log:     log.Sugar(),
		options: o,
	}, nil
}
