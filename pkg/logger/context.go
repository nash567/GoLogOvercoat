package logger

import (
	"context"

	"github.com/goLogOverCoat/pkg/logger/model"
)

func NewContextWithLogger(ctx context.Context, log model.Logger) context.Context {
	return context.WithValue(ctx, model.ContextKeyLogger, log)
}

//nolint:ireturn
func FromContext(ctx context.Context) model.Logger {
	logger, ok := ctx.Value(model.ContextKeyLogger).(model.Logger)
	if !ok || logger == nil {
		dLog := defaultLogger()
		if dLog != nil {
			dLog.Warn("logger instance not found in context")
		}
		return dLog
	}
	return logger
}
