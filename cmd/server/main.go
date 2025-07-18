package main

import (
	"context"
	"errors"
	"go-task-tracker/internal/http-server/handlers/memo"
	"go-task-tracker/internal/repository"
	"go-task-tracker/internal/service"
	"go-task-tracker/pkg/storage"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	log.Info("Starting the application")
	log.Debug("Debugging is enabled")

	godotenv.Load()
	postgresDsn := os.Getenv("POSTGRES_DSN")

	db, err := storage.NewPostgres(postgresDsn)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	httpServer := initHttpServer(log, db)
	ctx := context.Background()

	gracefulShutdown(ctx, log, httpServer)
}

func initHttpServer(log *slog.Logger, db *sqlx.DB) *http.Server {
	const serverAddress = ":8080"
	const readTimeout = 60
	const idleTimeout = 5

	memoRepo := repository.NewMemoRepository(db)
	memoService := service.NewMemoService(memoRepo)
	memoController := memo.NewMemoController(memoService)

	mux := http.NewServeMux()
	memoController.Register(mux, log)

	httpServer := &http.Server{
		ReadTimeout: readTimeout * time.Second,
		IdleTimeout: idleTimeout * time.Second,
		Addr:        serverAddress,
		Handler:     mux,
	}

	go func() {
		log.Info("Starting HTTP server", "address", serverAddress)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err.Error())
			os.Exit(1)
		}
	}()

	return httpServer
}

func gracefulShutdown(ctx context.Context, log *slog.Logger, server *http.Server) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit

	log.Info("Gracefully shutting down")
	server.Shutdown(ctx)
}
