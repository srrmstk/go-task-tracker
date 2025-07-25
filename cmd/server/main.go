package main

import (
	"context"
	"errors"
	"go-task-tracker/internal/http-server/handlers/auth"
	"go-task-tracker/internal/http-server/handlers/category"
	"go-task-tracker/internal/http-server/handlers/memo"
	jsonformatter "go-task-tracker/internal/http-server/middleware/json-formatter"
	"go-task-tracker/internal/repository"
	"go-task-tracker/internal/service"
	"go-task-tracker/pkg/storage"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	const readTimeout = 60 * time.Second
	const idleTimeout = 5 * time.Second

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(jsonformatter.JsonMiddleware)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := category.NewCategoryController(categoryService)

	memoRepo := repository.NewMemoRepository(db)
	memoService := service.NewMemoService(memoRepo, categoryRepo)
	memoController := memo.NewMemoController(memoService)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := auth.NewAuthController(authService)

	memoController.Register(r)
	categoryController.Register(r)
	authController.Register(r)

	httpServer := &http.Server{
		ReadTimeout: readTimeout,
		IdleTimeout: idleTimeout,
		Addr:        serverAddress,
		Handler:     r,
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
