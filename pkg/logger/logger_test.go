package logger_test

import (
	"errors"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/goLogOverCoat/pkg/logger"
	"github.com/goLogOverCoat/pkg/logger/model"
	"github.com/stretchr/testify/assert"
)

const (
	testMsgText    = "This is a test"
	testMsgText123 = "This is a test: 123"
)

var errLogger = errors.New("logger error")

func makeTestLogger() (*logger.SlogLogger, *strings.Builder) {
	output := new(strings.Builder)
	slogLogger := logger.NewSlogLogger(&model.Config{Output: output, IncludeSource: false, Level: model.InfoLevel.String()})
	return slogLogger, output
}

func outputMustMatch(t *testing.T, fname string, output string, patterns []string) {
	t.Helper()
	for _, p := range patterns {
		re := regexp.MustCompile(p)
		if !re.MatchString(output) {
			t.Fatalf(
				"failure in %s: output does not match pattern\n\tpattern: %v\n\toutput: %v",
				fname, p, output)
		}
	}
}

func TestNewSlogLogger(t *testing.T) {
	t.Parallel()
	type args struct {
		config *model.Config
	}
	tests := []struct {
		name string
		args args
		want *logger.SlogLogger
	}{
		{
			name: "slog with level debug",
			args: args{
				config: &model.Config{
					Level: model.DebugLevel.String(),
				},
			},
			want: &logger.SlogLogger{},
		},
		{
			name: "slog with level info",
			args: args{
				config: &model.Config{
					Level: model.InfoLevel.String(),
				},
			},
			want: &logger.SlogLogger{},
		},
		{
			name: "slog with level warn",
			args: args{
				config: &model.Config{
					Level: model.WarnLevel.String(),
				},
			},
			want: &logger.SlogLogger{},
		},
		{
			name: "slog with level error",
			args: args{
				config: &model.Config{
					Level: model.ErrorLevel.String(),
				},
			},
			want: &logger.SlogLogger{},
		},
		{
			name: "slog with level fatal",
			args: args{
				config: &model.Config{
					Level: model.FatalLevel.String(),
				},
			},
			want: &logger.SlogLogger{},
		},
	}
	for _, tC := range tests {
		tt := tC
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			slog := logger.NewSlogLogger(tt.args.config)
			if reflect.TypeOf(slog) != reflect.TypeOf(tt.want) {
				t.Errorf("wrong logger type: %v", slog)
			}
		})
	}
}

func TestSlogLogger_Debug(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()
	slogLogger.SetLevel(model.DebugLevel)

	slogLogger.Debug("This is a test")
	outputMustMatch(t, "SlogLogger.Debug", output.String(), []string{
		testString(model.DebugLevel, testMsgText),
	})
}

func TestSlogLogger_Debugf(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()
	slogLogger.SetLevel(model.DebugLevel)

	slogLogger.Debugf("This is a test: %d", 123)
	outputMustMatch(t, "SlogLogger.Debugf", output.String(), []string{
		testString(model.DebugLevel, testMsgText123),
	})
}

func TestSlogLogger_Info(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	slogLogger.Info("This is a test")
	outputMustMatch(t, "SlogLogger.Info", output.String(), []string{
		testString(model.InfoLevel, testMsgText),
	})
}

func TestSlogLogger_Infof(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	slogLogger.Infof("This is a test: %d", 123)
	outputMustMatch(t, "SlogLogger.Infof", output.String(), []string{
		testString(model.InfoLevel, testMsgText123),
	})
}

func TestSlogLogger_Warn(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	slogLogger.Warn("This is a test")
	outputMustMatch(t, "SlogLogger.Warn", output.String(), []string{
		testString(model.WarnLevel, testMsgText),
	})
}

func TestSlogLogger_Fatal(t *testing.T) {
	if os.Getenv("CALL_LOG_FATAL") == "1" {
		config := model.Config{Output: os.Stdout}
		slog := logger.NewSlogLogger(&config)

		slog.Fatal("This is a test")
	}
	t.Parallel()
	outb := new(strings.Builder)
	cmString := os.Args[0]
	cmd := exec.Command(cmString, "-test.run=TestSlogLogger_Fatal")
	cmd.Env = append(os.Environ(), "CALL_LOG_FATAL=1")
	cmd.Stdout = outb

	errExit := new(exec.ExitError)
	if err := cmd.Run(); errors.As(err, &errExit) && !errExit.Success() {
		outputMustMatch(t, "SlogLogger.Fatal", outb.String(), []string{
			testString(model.FatalLevel, testMsgText),
		})
		return
	}

	t.Fatal("failure in ZapLogger.Fatal: calling did not result in os.Exit(1)")
}

func TestSlogLogger_Fatalf(t *testing.T) {
	if os.Getenv("CALL_LOG_FATAL") == "1" {
		config := model.Config{Output: os.Stdout}
		slog := logger.NewSlogLogger(&config)

		slog.Fatalf("This is a test: %d", 123)
	}
	t.Parallel()
	outb := new(strings.Builder)
	cmString := os.Args[0]
	cmd := exec.Command(cmString, "-test.run=TestSlogLogger_Fatalf")
	cmd.Env = append(os.Environ(), "CALL_LOG_FATAL=1")
	cmd.Stdout = outb

	errExit := new(exec.ExitError)
	if err := cmd.Run(); errors.As(err, &errExit) && !errExit.Success() {
		outputMustMatch(t, "SlogLogger.Fatal", outb.String(), []string{
			testString(model.FatalLevel, testMsgText123),
		})
		return
	}

	t.Fatal("failure in ZapLogger.Fatal: calling did not result in os.Exit(1)")
}

func TestSlogLogger_Warnf(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	slogLogger.Warnf("This is a test: %d", 123)
	outputMustMatch(t, "SlogLogger.Warnf", output.String(), []string{
		testString(model.WarnLevel, testMsgText123),
	})
}

func TestSlogLogger_Error(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	slogLogger.Error("This is a test")
	outputMustMatch(t, "SlogLogger.Error", output.String(), []string{
		testString(model.ErrorLevel, testMsgText),
	})
}

func TestSlogLogger_Errorf(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	slogLogger.Errorf("This is a test: %d", 123)
	outputMustMatch(t, "SlogLogger.Errorf", output.String(), []string{
		testString(model.ErrorLevel, testMsgText123),
	})
}

func TestSlogLogger_SetLevel(t *testing.T) {
	t.Parallel()
	slogLogger, output := makeTestLogger()

	var setter model.LevelSetter = slogLogger

	setter.SetLevel(model.InfoLevel)

	slogLogger.Info("info msg")
	slogLogger.Debug("debug msg")
	if strings.Contains(output.String(), "debug msg") {
		t.Error("unexpected 'debug msg' at log records")
	}

	slogLogger, output = makeTestLogger()
	setter = slogLogger

	setter.SetLevel(model.DebugLevel)

	slogLogger.Debug("debug msg")
	if !strings.Contains(output.String(), "debug msg") {
		t.Error("expected 'debug msg' at log records")
	}
}

func TestSlogLogger_WithField(t *testing.T) {
	t.Parallel()
	var (
		l      model.Logger
		output *strings.Builder
	)
	l, output = makeTestLogger()
	l = l.WithField("key", "demo")
	l.Error("This is a test")
	outputMustMatch(t, "SlogLogger.Error", output.String(), []string{
		testString(model.ErrorLevel, testMsgText, "key", "demo"),
	})
}

func TestSlogLogger_WithFields(t *testing.T) {
	t.Parallel()
	var (
		l      model.Logger
		output *strings.Builder
	)
	l, output = makeTestLogger()
	l = l.WithFields(model.Fields{"key": "demo"})
	l.Error("This is a test")
	outputMustMatch(t, "SlogLogger.Error", output.String(), []string{
		testString(model.ErrorLevel, testMsgText, "key", "demo"),
	})
}

func TestSlogLogger_WithError(t *testing.T) {
	t.Parallel()
	var (
		log    model.Logger
		output *strings.Builder
	)
	log, output = makeTestLogger()
	log = log.WithError(errLogger)
	log.Error("This is a test")
	outputMustMatch(t, "SlogLogger.Error", output.String(), []string{
		testString(model.ErrorLevel, testMsgText, "error", errLogger.Error()),
	})
}

func TestSlogLogger_GetLevel(t *testing.T) {
	t.Parallel()
	slogLogger, _ := makeTestLogger()
	var setter model.LevelSetter = slogLogger
	setter.SetLevel(model.DebugLevel)

	assert.Equal(t, model.DebugLevel, setter.GetLevel())
}

func testString(level model.Level, msg string, kv ...string) string {
	kvStr := make([]string, 0, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		val := ""
		if i+1 < len(kv) {
			val = kv[i+1]
		}
		kvStr = append(kvStr, `"`+kv[i]+`":"`+val+`"`)
	}
	resp := `{"time":".*","level":"` + level.String() + `","msg":"` + msg + `"`

	if len(kvStr) > 0 {
		resp += "," + strings.Join(kvStr, ",")
	}
	return resp + `}`
}
