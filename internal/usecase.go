package main

import (
	"context"
	"log/slog"
)

type UseCase struct {
	Repo Repo
}

func (u *UseCase) GetInfo(ctx context.Context) (string, error) {
	slog.InfoContext(ctx, "In useCase")
	inf, err := u.Repo.GetInfo(ctx)
	if err != nil {
		return "", err
	}
	slog.InfoContext(ctx, "Got Info From Repo")
	return inf, nil
}
