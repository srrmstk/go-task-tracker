package main

import (
	"context"
	"errors"
	"go-task-tracker/internal/helpers"
	"go-task-tracker/internal/http-server/handlers/auth"
	"go-task-tracker/internal/http-server/handlers/category"
	"go-task-tracker/internal/http-server/handlers/memo"
	"go-task-tracker/internal/http-server/middleware"
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
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	log.Info("Starting the application")
	log.Debug("Debugging is enabled")
	slog.SetDefault(log)

	godotenv.Load()
	postgresDsn := os.Getenv("POSTGRES_DSN")
	redisDsn := os.Getenv("REDIS_DSN")

	db, err := storage.NewPostgres(postgresDsn)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	rdb, err := storage.NewRedis(redisDsn)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	defer rdb.Close()

	httpServer := initHttpServer(db, rdb)
	ctx := context.Background()

	gracefulShutdown(ctx, httpServer)
}

func initHttpServer(db *sqlx.DB, rdb *redis.Client) *http.Server {
	const serverAddress = ":8080"
	const readTimeout = 60 * time.Second
	const idleTimeout = 5 * time.Second

	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(middleware.JsonMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		helpers.JsonError(w, "Not found", http.StatusNotFound)
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		helpers.JsonError(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	emailService := service.NewEmailService()

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := category.NewCategoryController(categoryService)

	memoRepo := repository.NewMemoRepository(db)
	memoService := service.NewMemoService(memoRepo, categoryRepo)
	memoController := memo.NewMemoController(memoService)

	authRepo := repository.NewAuthRepository(db, rdb)
	authService := service.NewAuthService(authRepo, emailService)
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
		slog.Info("Starting HTTP server", "address", serverAddress)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(err.Error())
			os.Exit(1)
		}
	}()

	return httpServer
}

func gracefulShutdown(ctx context.Context, server *http.Server) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit

	slog.Info("Gracefully shutting down")
	server.Shutdown(ctx)
}

// TODO: add redis queue
