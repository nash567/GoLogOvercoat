package model

import (
	"io"
	"log/slog"
)

// Config is a logging configuration.
type Config struct {
	// Level is the lowest level of log message that should be emitted. Any log
	// messages logged at the specified level or any level more severe will be
	// emitted. The default level is INFO.
	Level string `default:"INFO"`

	// Output is the destination for log messages. By default, it is os.Stdout.
	Output io.Writer

	// IncludeSource specifies whether to add source in the output. Default is false.
	IncludeSource bool
}

func (c *Config) GetLevel() Level {
	return ParseLevel(c.Level)
}

func (c *Config) GetSlogLevel() slog.Level {
	return c.GetLevel().SlogLevel()
}
