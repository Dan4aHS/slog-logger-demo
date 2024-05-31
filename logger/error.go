package logger

import (
	"context"
	"errors"
	"runtime"
)

type ErrorWithLogContext struct {
	next error
	ctx  LogContext
}

type ErrorInfo struct {
	InternalCode int
	StatusCode   int
	Message      string
	Source       uintptr
	FunctionName string
}

func (e *ErrorWithLogContext) Error() string {
	return e.next.Error()
}

func LogError(ctx context.Context, err error, internalCode int, statusCode int) error {
	pc, _, _, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	ei := ErrorInfo{
		Source:       pc,
		InternalCode: internalCode,
		StatusCode:   statusCode,
		Message:      err.Error(),
		FunctionName: functionName,
	}
	c := LogContext{
		ErrorInfo: ei,
	}
	if x, ok := ctx.Value(key).(LogContext); ok {
		x.ErrorInfo = ei
		c = x
	}
	return &ErrorWithLogContext{next: err, ctx: c}
}

func ErrorCtx(ctx context.Context, err error) context.Context {
	var ctxErr *ErrorWithLogContext
	if errors.As(err, &ctxErr) {
		return context.WithValue(ctx, key, ctxErr.ctx)
	}
	return ctx
}
