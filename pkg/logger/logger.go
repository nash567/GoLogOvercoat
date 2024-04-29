package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/goLogOverCoat/pkg/logger/model"
)

// SlogLogger is the default implementation of Logger. It is backed by the slog logging package.
type SlogLogger struct {
	entry *slog.Logger
	cfg   *model.Config
	level *slog.LevelVar
}

func NewSlogLogger(config *model.Config) *SlogLogger {
	loggingLevel := new(slog.LevelVar)
	loggingLevel.Set(config.GetSlogLevel())

	handler := slog.NewJSONHandler(config.Output, &slog.HandlerOptions{
		AddSource:   config.IncludeSource,
		Level:       loggingLevel,
		ReplaceAttr: replaceAttribute,
	})
	s := &SlogLogger{
		entry: slog.New(handler),
		cfg:   config,
		level: loggingLevel,
	}

	// output from the log package's default Logger (as with log.Print, etc.) will be logged using slog Handler
	slog.SetDefault(s.entry)
	return s
}

// func to map the custom log levels to their respective labels
// e.g.-> slog doesn't have FATAL level.
func replaceAttribute(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level, ok := a.Value.Any().(slog.Level)
		if !ok {
			return a
		}
		levelLabel, exists := getCustomLevelMap()[level]
		if !exists {
			levelLabel = level.String()
		}
		a.Value = slog.StringValue(levelLabel)
	}
	return a
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

func (log *SlogLogger) Fatal(msg string) {
	log.entry.Log(context.Background(), model.LevelFatal, msg)
	os.Exit(1)
}

func (log *SlogLogger) Fatalf(format string, args ...interface{}) {
	log.entry.Log(context.Background(), model.LevelFatal, fmt.Sprintf(format, args...))
	os.Exit(1)
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

func (log *SlogLogger) SetLevel(lvl model.Level) {
	log.level.Set(lvl.SlogLevel())
}

func (log *SlogLogger) GetLevel() model.Level {
	return model.ParseLevel(log.level.Level().String())
}

func getDefaultLogger() *SlogLogger {
	defaultLogger := slog.Default()
	levelVar := new(slog.LevelVar)
	return &SlogLogger{
		entry: defaultLogger,
		level: levelVar,
	}
}

// map contains custom slog levels.
func getCustomLevelMap() map[slog.Leveler]string {
	return map[slog.Leveler]string{
		model.LevelFatal: "FATAL",
	}
}
