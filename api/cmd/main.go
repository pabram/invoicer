package main

import (
	"api/invoice"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

const dbsource = "postgresql://postgres:postgres@db:5432/invoice?sslmode=disable"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
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

	flag.Parse()
	ctx := context.Background()

	repository := invoice.NewRepository(db, logger)
	srv := invoice.NewService(repository, logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := invoice.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := invoice.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
