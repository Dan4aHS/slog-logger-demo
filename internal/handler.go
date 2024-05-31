package main

import (
	"bwg_logger/logger"
	"log/slog"
	"net/http"
	"time"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	repo := Repo{Info: "some_info"}
	uc := UseCase{Repo: repo}
	ctx = logger.WithLogRequest(ctx, r)
	slog.InfoContext(ctx, r.URL.String())
	ctx = logger.WithLogTime(ctx, time.Now())
	_, err := uc.GetInfo(ctx)
	slog.DebugContext(ctx, "SOME DEBUG MSG")
	if err != nil {
		slog.ErrorContext(logger.ErrorCtx(ctx, err), err.Error())
	}
	ctx = logger.WithLogAttributes(ctx, "status_code", http.StatusOK)
	slog.InfoContext(ctx, "ResponseOK")
	slog.Log(ctx, logger.LevelFatal, "FatalError")
}
