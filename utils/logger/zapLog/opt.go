package zapLog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	options struct {
		timeEncoder zapcore.TimeEncoder
		zapOptions  []zap.Option
	}

	Option interface {
		Apply(o *options)
	}
)

func defaultOptions() options {
	return options{
		timeEncoder: zapcore.RFC3339TimeEncoder,
	}
}

type withTimeEncoder struct{ zapcore.TimeEncoder }

func (w withTimeEncoder) Apply(o *options) {
	o.timeEncoder = w.TimeEncoder
}

func WithTimeEncoder(timeEncoder zapcore.TimeEncoder) Option {
	return &withTimeEncoder{timeEncoder}
}

type withZapOption struct{ zap.Option }

func (w withZapOption) Apply(o *options) {
	o.zapOptions = append(o.zapOptions, w.Option)
}

func WithZapOption(zapOpt zap.Option) Option {
	return &withZapOption{zapOpt}
}
