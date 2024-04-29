package logger

import (
	"context"

	"github.com/goLogOverCoat/pkg/logger/model"
)

func NewContextWithLogger(ctx context.Context, log model.Logger) context.Context {
	return context.WithValue(ctx, model.ContextKeyLogger, log)
}

func FromContext(ctx context.Context) *SlogLogger {
	logger, ok := ctx.Value(model.ContextKeyLogger).(*SlogLogger)
	if !ok || logger == nil {
		dLog := getDefaultLogger()
		if dLog != nil {
			dLog.Warn("logger instance not found in context")
		}
		return dLog
	}
	return logger
}
