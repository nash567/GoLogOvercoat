package logger

import (
	"fmt"
	"log"
	"log/slog"
	"sync"

	"github.com/goLogOverCoat/pkg/logger/model"
)

//nolint:gochecknoglobals // to be used to return default logger when not set in context
var (
	mux        = &sync.Mutex{}
	defaultLog *SlogLogger
)

func defaultLogger() *SlogLogger {
	mux.Lock()
	defer mux.Unlock()
	return defaultLog
}

// SlogLogger is the default implementation of Logger. It is backed by the slog logging package.
type SlogLogger struct {
	entry *slog.Logger
	cfg   *model.Config
	level *slog.LevelVar
}

var _ model.Logger = &SlogLogger{}

func NewSlogLogger(config *model.Config) *SlogLogger {
	loggingLevel := new(slog.LevelVar)
	loggingLevel.Set(config.GetSlogLevel())
	handler := slog.NewJSONHandler(config.Output, &slog.HandlerOptions{
		AddSource: config.IncludeSource,
		Level:     loggingLevel,
	})
	s := &SlogLogger{
		entry: slog.New(handler),
		cfg:   config,
		level: loggingLevel,
	}
	// output from the log package's default Logger (as with log.Print, etc.) will be logged using slog Handler
	slog.SetDefault(s.entry)

	mux.Lock()
	defer mux.Unlock()
	defaultLog = s
	return s
}

func (log *SlogLogger) Debug(msg string) {

	log.entry.Debug(msg)

}

func (log *SlogLogger) Debugf(format string, args ...interface{}) {

	log.entry.Debug(fmt.Sprintf(format, args...))

}

func (log *SlogLogger) Info(msg string) {

	log.entry.Info(msg)

}

func (log *SlogLogger) Infof(format string, args ...interface{}) {

	log.entry.Info(fmt.Sprintf(format, args...))

}

func (log *SlogLogger) Warn(msg string) {

	log.entry.Warn(msg)

}

func (log *SlogLogger) Warnf(format string, args ...interface{}) {

	log.entry.Warn(fmt.Sprintf(format, args...))

}

func (log *SlogLogger) Error(msg string) {

	log.entry.Error(msg)

}

func (log *SlogLogger) Errorf(format string, args ...interface{}) {

	log.entry.Error(fmt.Sprintf(format, args...))

}

//nolint:ireturn // implements model.Logger interface
func (log *SlogLogger) WithField(key string, value interface{}) model.Logger {
	return &SlogLogger{
		entry: log.entry.With(key, value),
		level: log.level,
	}
}

//nolint:ireturn // implements model.Logger interface
func (log *SlogLogger) WithFields(fields model.Fields) model.Logger {
	sFields := make([]any, 0)
	for key, value := range fields {
		sFields = append(sFields, key, value)
	}
	return &SlogLogger{
		entry: log.entry.With(sFields...),
		level: log.level,
	}
}

//nolint:ireturn // implements model.Logger interface
func (log *SlogLogger) WithError(err error) model.Logger {
	return &SlogLogger{
		entry: log.entry.With("error", err),
		level: log.level,
	}
}

// ToStdLogger creates a logger that matches std library.
func (log *SlogLogger) ToStdLogger() *log.Logger {
	handler := log.entry.Handler()
	return slog.NewLogLogger(handler, log.cfg.GetSlogLevel())
}

func (log *SlogLogger) SetLevel(lvl model.Level) error {
	mux.Lock()
	defer mux.Unlock()
	log.level.Set(lvl.SlogLevel())
	return nil
}

func (log *SlogLogger) GetLevel() model.Level {
	return model.ParseLevel(log.level.Level().String())
}

//
// func getDefaultLogger() model.Logger {
//	return NewSlogLogger(&model.Config{
//		Level: "INFO",
//	})
//}
