package model

import (
	"log/slog"
	"strings"
)

const (
	//slog does not have fatal level creating a custom fatal level
	LevelFatal = slog.Level(13)
)

// A Level is a level of severity for a log message.
type Level uint8

const (

	// DebugLevel causes a logger to emit messages logged at "DEBUG" level or more
	// severe. It is typically only enabled when debugging or during development,
	// and usually results in very verbose logging output.
	DebugLevel Level = iota + 1

	// InfoLevel causes a logger to emit messages logged at "INFO" level or more
	// severe. It is typically used for general operational entries about what's
	// going on inside an application.
	InfoLevel

	// WarnLevel causes a logger to emit messages logged at "WARN" level or more
	// severe. It is typically used for non-critical entries that deserve attention.
	WarnLevel

	// ErrorLevel causes a logger to emit messages logged at "ERROR" level or more
	// severe. It is typically used for errors that should definitely be noted, and
	// is commonly used for hooks to send errors to an error tracking service.
	ErrorLevel

	// FatalLevel causes a logger to emit messages logged at "FATAL" level or more
	// severe. Messages logged at this level cause a logger to log the message and
	// then call os.Exit(1).
	FatalLevel
)

// String implements fmt.Stringer for Level.
func (l Level) String() string {
	switch l {
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	case FatalLevel:
		return "FATAL"
	default:
		return "INFO"
	}
}

// ParseLevel converts log level string to level constant
//
//	if the wrong string received it returns info level.
func ParseLevel(logLevel string) Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

func (l Level) SlogLevel() slog.Level {
	switch l {
	case WarnLevel:
		return slog.LevelWarn
	case InfoLevel:
		return slog.LevelInfo
	case DebugLevel:
		return slog.LevelDebug
	case ErrorLevel:
		return slog.LevelError
	case FatalLevel:
		return LevelFatal
	default:
		return slog.LevelInfo
	}
}
