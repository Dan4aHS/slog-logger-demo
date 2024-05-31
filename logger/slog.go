package logger

import (
	"context"
	console "github.com/phsym/console-slog"
	"log/slog"
	"os"
)

type ILogger interface {
	LogFatal(ctx context.Context, msg string)
	LogError(ctx context.Context, msg string)
	LogWarn(ctx context.Context, msg string)
	LogInfo(ctx context.Context, msg string)
	LogDebug(ctx context.Context, msg string)
}

type Logger struct {
}

func InitLogging() *slog.Handler {
	handler := slog.Handler(console.NewHandler(os.Stdout, &console.HandlerOptions{AddSource: true, Level: slog.LevelDebug}))
	handler = NewLogMiddleware(handler)
	return &handler
	//slog.SetDefault(slog.New(handler))
}

// Esli komu to nado logger v ctx ili logger vo vse structuri :clownmask:
func (l *Logger) Fatal(ctx context.Context, msg string) {
	slog.Log(ctx, LevelFatal, msg)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	slog.Log(ctx, LevelError, msg)
}

func (l *Logger) Warn(ctx context.Context, msg string) {
	slog.Log(ctx, LevelWarn, msg)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	slog.Log(ctx, LevelInfo, msg)
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	slog.Log(ctx, LevelDebug, msg)
}
