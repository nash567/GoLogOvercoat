package model

import "log"

// A Logger provides methods for logging messages.
type Logger interface {
	// Debug emits a "DEBUG" level log message.
	Debug(msg string)

	// Debugf emits a "DEBUG" level log message. Additional arguments are applied
	// to format as string formatting parameters.
	Debugf(format string, args ...interface{})

	// Info emits an "INFO" level log message.
	Info(msg string)

	// Infof emits an "INFO" level log message. Additional arguments are applied
	// to format as string formatting parameters.
	Infof(format string, args ...interface{})

	// Warn emits a "WARN" level log message.
	Warn(msg string)

	// Warnf emits a "WARN" level log message. Additional arguments are applied
	// to format as string formatting parameters.
	Warnf(format string, args ...interface{})

	// Error emits an "ERROR" level log message.
	Error(msg string)

	// Errorf emits an "ERROR" level log message. Additional arguments are applied
	// to format as string formatting parameters.
	Errorf(format string, args ...interface{})

	// WithField adds a field to the logger and returns a new Logger.
	WithField(key string, value interface{}) Logger

	// WithFields adds multiple fields to the logger and returns a new Logger.
	WithFields(fields Fields) Logger

	// WithError adds a field called "error" to the logger and returns a new Logger.
	WithError(err error) Logger

	// ToStdLogger makes logger that satisfies std logger interface. All msg are logged at Info level.
	ToStdLogger() *log.Logger
}

//go:generate mockery --name=Logger --outpkg mocks

type LevelSetter interface {
	SetLevel(Level) error
	GetLevel() Level
}

//go:generate mockery --name=LevelSetter --outpkg mocks
