package main

import (
	"context"
	"errors"
	"go-task-tracker/internal/config"
	"go-task-tracker/internal/logger"
	"go-task-tracker/internal/storage"
	goLog "log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configPath := "config/local.yml"

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		goLog.Fatalf("%v\n", err)
	}

	log := logger.InitLogger(cfg.Env)

	log.Info("Starting the application", "env", cfg.Env)
	log.Debug("Debugging is enabled")

	db, err := storage.NewPostgres(cfg.Database.PostgresDSN)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	_ = db

	httpServer := initHttpServer(cfg, log)

	ctx := context.Background()

	gracefulShutdown(ctx, log, httpServer)
}

func initHttpServer(cfg *config.Config, log *slog.Logger) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	httpServer := &http.Server{
		ReadTimeout: cfg.HTTPServer.Timeout * time.Second,
		IdleTimeout: cfg.HTTPServer.IdleTimeout * time.Second,
		Addr:        cfg.HTTPServer.Address,
		Handler:     mux,
	}

	go func() {
		log.Info("Starting HTTP server", "address", cfg.HTTPServer.Address)
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
