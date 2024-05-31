package logger

import (
	"context"
	"log/slog"
	"time"
	"unsafe"
)

type LogMiddleware struct {
	next slog.Handler
}

func NewLogMiddleware(next slog.Handler) *LogMiddleware {
	return &LogMiddleware{next: next}
}

func (l *LogMiddleware) Enabled(ctx context.Context, rec slog.Level) bool {
	return l.next.Enabled(ctx, rec)
}

func (l *LogMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if rec.Level == LevelDebug {
		return l.next.Handle(ctx, rec)
	}
	if c, ok := ctx.Value(key).(LogContext); ok {
		if c.Public != "" {
			rec.Add("public", c.Public)
		}
		if c.Access != "" {
			rec.Add("access", c.Access)
		}
		if c.Signature != "" {
			rec.Add("signature", c.Signature)
		}
		if c.RequestBodyString != "" {
			rec.Add("requestBody", c.RequestBodyString)
		}
		if c.ResponseBodyString != "" {
			rec.Add("responseBody", c.ResponseBodyString)
		}
		if c.StatusCode != 0 {
			rec.Add("statusCode", c.StatusCode)
		}
		if !c.StartTime.IsZero() {
			rec.Add("duration", time.Since(c.StartTime))
		}
		if rec.Level == LevelError && c.ErrorInfo.Source != uintptr(unsafe.Pointer(nil)) {
			rec.PC = c.ErrorInfo.Source
			rec.Add("function", c.ErrorInfo.FunctionName)
			rec.Add("status_code", c.ErrorInfo.StatusCode)
			rec.Add("internal_code", c.ErrorInfo.InternalCode)
		}
		//fmt.Println(c.Attributes)
		if c.Attributes != nil {
			for k, v := range c.Attributes {
				rec.Add(k, v)
			}
		}
	}

	return l.next.Handle(ctx, rec)
}

func (l *LogMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LogMiddleware{next: l.next.WithAttrs(attrs)}
}

func (l *LogMiddleware) WithGroup(name string) slog.Handler {
	return &LogMiddleware{next: l.next.WithGroup(name)}
}
