package main

import (
	"api/invoice"
	"context"
	"database/sql"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

const dbsource = "postgresql://postgres:postgres@db:5432/invoice?sslmode=disable"

func main() {
	var logger log.Logger

	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger, "service", "invoice", "time:", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "service starting")
	defer level.Info(logger).Log("msg", "service ending")

	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", dbsource)
	if err != nil {
		level.Error(logger).Log("database connection error", err)
		os.Exit(-1)
	}

	repository := invoice.NewRepository(db, logger)
	srv := invoice.NewService(repository, logger)

	ctx := context.Background()

	router := invoice.NewRouter(ctx, srv)
	router.Run(":5000")
}
