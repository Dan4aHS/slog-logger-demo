package logger

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

type LogContext struct {
	UserID             int
	Public             string
	Access             string
	Signature          string
	StatusCode         int
	StartTime          time.Time
	RequestBodyString  string
	ResponseBodyString string
	Attributes         map[string]any
	// ErrorSource        uintptr
	ErrorInfo ErrorInfo
}

type keyType int

const key = keyType(0)

func WithLogUserID(ctx context.Context, userID int) context.Context {
	if c, ok := ctx.Value(key).(LogContext); ok {
		c.UserID = userID
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, LogContext{UserID: userID})
}

func WithLogRequest(ctx context.Context, r *http.Request) context.Context {
	var (
		access, public, signature, reqBody string
		headers                            = r.Header
		body                               = r.Body
	)

	access = headers.Get("Access")
	public = headers.Get("Public")
	signature = headers.Get("Signature")
	if body != nil {
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			bodyBytes = []byte{}
		}
		reqBody = string(bodyBytes)
	}

	if len(access) > 8 {
		access = access[:8] + strings.Repeat("*", len(access)-8)
	}
	if len(signature) > 4 {
		signature = signature[:4] + strings.Repeat("*", len(signature)-4)
	}

	if c, ok := ctx.Value(key).(LogContext); ok {
		c.Public = public
		c.Access = access
		c.Signature = signature
		c.RequestBodyString = reqBody
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, LogContext{
		Public:            public,
		Access:            access,
		Signature:         signature,
		RequestBodyString: reqBody,
	})
}

func WithLogResponse(ctx context.Context, response http.Response) context.Context {
	var (
		statusCode = response.StatusCode
		body       = response.Body
		resBody    string
	)

	if body != nil {
		bodyBytes, err := io.ReadAll(body)
		if err != nil {
			bodyBytes = []byte{}
		}
		resBody = string(bodyBytes)
	}

	if c, ok := ctx.Value(key).(LogContext); ok {
		c.StatusCode = statusCode
		c.ResponseBodyString = resBody
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, LogContext{
		StatusCode:         statusCode,
		ResponseBodyString: resBody,
	})
}

func WithLogTime(ctx context.Context, startTime time.Time) context.Context {
	if c, ok := ctx.Value(key).(LogContext); ok {
		c.StartTime = startTime
		return context.WithValue(ctx, key, c)
	}
	return context.WithValue(ctx, key, LogContext{StartTime: startTime})
}

func WithLogAttributes(ctx context.Context, keym string, value any) context.Context {
	if c, ok := ctx.Value(key).(LogContext); ok {
		if c.Attributes == nil {
			c.Attributes = make(map[string]any)
		}
		c.Attributes[keym] = value
		//fmt.Println(c.Attributes)
		return context.WithValue(ctx, key, c)
	}
	attrs := make(map[string]any)
	attrs[keym] = value
	//fmt.Println(attrs)
	return context.WithValue(ctx, key, LogContext{Attributes: attrs})
}
