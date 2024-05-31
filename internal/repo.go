package main

import (
	"bwg_logger/logger"
	"context"
	"errors"
	"log/slog"
)

type Repo struct {
	Info string
}

func (r *Repo) GetInfo(ctx context.Context) (string, error) {
	var err error
	slog.InfoContext(ctx, "In repo")
	err = errors.New("error from repo")
	if err != nil {
		return "", logger.LogError(ctx, err, 300, 200)
	}
	slog.InfoContext(ctx, "Success Getting Info from Repo")
	return r.Info, nil
}
