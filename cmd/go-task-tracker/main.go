package main

import (
	"go-task-tracker/internal/config"
	"go-task-tracker/internal/logger"
	"go-task-tracker/internal/storage"
	goLog "log"
	"os"
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

	log.Info(db.DriverName())
	// TODO: init router - net/http

	// http.HandleFunc("/tasks", handleTasks)
	//http.ListenAndServe(":8000", nil)
}
