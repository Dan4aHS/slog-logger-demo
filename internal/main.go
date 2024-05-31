package main

import (
	"bwg_logger/logger"
	"context"
	"log/slog"
	"net/http"
)

func main() {
	slog.SetDefault(slog.New(*logger.InitLogging()))

	http.HandleFunc("/info", InfoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Log(context.Background(), logger.LevelFatal, "Could not start server", err)
	}
	//ctx = logger.WithLogUserID(ctx, 111)
	//req := http.Request{
	//	Header: http.Header{
	//		"Public":    []string{"qwe123qwirqwe2134ue"},
	//		"Access":    []string{"tqyfrguovnfweqwere"},
	//		"Signature": []string{"qwkbgyvweqiowfewqpru"},
	//	},
	//	Body: nil,
	//}

}
