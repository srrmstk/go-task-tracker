package main

import (
	"go-task-tracker/internal/config"
	"go-task-tracker/internal/logger"
	"os"
)

func main() {
	configPath := "config/local.yml"

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		os.Exit(1)
	}

	log := logger.InitLogger(cfg.Env)

	log.Info("Starting the application", "env", cfg.Env)
	log.Debug("Debugging is enabled")

	// TODO: init storage - postgres + redis
	// TODO: init router - net/http

	// http.HandleFunc("/tasks", handleTasks)
	// http.ListenAndServe(":8000", nil)
}
