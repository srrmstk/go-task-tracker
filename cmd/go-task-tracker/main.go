package main

import (
	"fmt"
	"go-task-tracker/internal/config"
	"os"
)

func main() {
	configPath := "config/local.yml"

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(cfg)

	// TODO: init logger - log/slog
	// TODO: init storage - postgres + redis
	// TODO: init router - net/http

	// http.HandleFunc("/tasks", handleTasks)
	// http.ListenAndServe(":8000", nil)
}
