package logger_test

import (
	"context"
	"testing"

	"github.com/goLogOverCoat/pkg/logger"
	"github.com/goLogOverCoat/pkg/logger/model"
	logMocks "github.com/goLogOverCoat/pkg/logger/model/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFromContext(t *testing.T) {
	t.Parallel()
	ctxLog := logger.NewSlogLogger(&model.Config{})
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		args       args
		wantCtxLog bool
	}{
		{
			name: "logger set in context",
			args: args{
				ctx: context.WithValue(context.Background(), model.ContextKeyLogger, ctxLog),
			},
			wantCtxLog: true,
		},
	}
	for _, tC := range tests {
		tt := tC
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := logger.FromContext(tt.args.ctx)
			if tt.wantCtxLog {
				assert.Same(t, ctxLog, got)
				return
			}
			assert.NotSame(t, ctxLog, got)
		})
	}
}

func TestNewContextWithLogger(t *testing.T) {
	t.Parallel()
	ctxLog := &logMocks.MockLogger{}
	type args struct {
		ctx context.Context
		log model.Logger
	}
	tests := []struct {
		name string
		args args
		want model.Logger
	}{
		{
			name: "positive case",
			args: args{
				ctx: context.Background(),
				log: ctxLog,
			},
			want: ctxLog,
		},
	}
	for _, tC := range tests {
		tt := tC
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotCtx := logger.NewContextWithLogger(tt.args.ctx, tt.args.log)
			assert.NotNil(t, gotCtx)
			got := gotCtx.Value(model.ContextKeyLogger)
			assert.NotNil(t, got)
			assert.Same(t, tt.want, got)
		})
	}
}
