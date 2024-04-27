package model_test

import (
	"log/slog"
	"testing"

	"github.com/goLogOverCoat/pkg/logger/model"
)

func TestParseLevel(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc string
		sent string
		want model.Level
	}{
		{desc: "empty string", sent: "", want: model.InfoLevel},
		{desc: "incorrect string", sent: "test", want: model.InfoLevel},
		{desc: "info string", sent: "info", want: model.InfoLevel},
		{desc: "warn string", sent: "warn", want: model.WarnLevel},
		{desc: "error string", sent: "error", want: model.ErrorLevel},
		{desc: "debug string", sent: "debug", want: model.DebugLevel},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			lvl := model.ParseLevel(tC.sent)
			if lvl != tC.want {
				t.Error("unexpected level")
			}
		})
	}
}

func TestLevelString(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input model.Level
		desc  string
		want  string
	}{
		{input: model.ErrorLevel, desc: "error level", want: "ERROR"},
		{input: model.WarnLevel, desc: "warn level", want: "WARN"},
		{input: model.InfoLevel, desc: "info level", want: "INFO"},
		{input: model.DebugLevel, desc: "debug level", want: "DEBUG"},
		{input: model.Level(100), desc: "default level", want: "INFO"},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			if tC.input.String() != tC.want {
				t.Error("unexpected level")
			}
		})
	}
}

func TestSlogLevel(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input model.Level
		desc  string
		want  slog.Level
	}{
		{input: model.ErrorLevel, desc: "error level", want: slog.LevelError},
		{input: model.WarnLevel, desc: "warn level", want: slog.LevelWarn},
		{input: model.InfoLevel, desc: "info level", want: slog.LevelInfo},
		{input: model.DebugLevel, desc: "debug level", want: slog.LevelDebug},
		{input: model.Level(100), desc: "default level", want: slog.LevelInfo},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			if tC.input.SlogLevel() != tC.want {
				t.Error("unexpected level")
			}
		})
	}
}

func TestLevel_Enabled(t *testing.T) {
	type args struct {
		lvl model.Level
	}
	tests := []struct {
		name string
		l    model.Level
		args args
		want bool
	}{
		{
			name: "Debug logs are enabled",
			l:    model.DebugLevel,
			args: args{
				lvl: model.InfoLevel,
			},
			want: true,
		},
		{
			name: "Debug logs are disabled",
			l:    model.InfoLevel,
			args: args{
				lvl: model.DebugLevel,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Enabled(tt.args.lvl); got != tt.want {
				t.Errorf("Enabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
