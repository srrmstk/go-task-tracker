run:
	@echo "Starting server"
	go run cmd/go-task-tracker/main.go

compose-dev:
	@echo "Starting docker-compose"
	docker-compose -f docker-compose.dev.yml up

compose-dev-detached:
	@echo "Starting docker-compose-build"
	docker-compose -f docker-compose.dev.yml up -d

compose-dev-down:
	@echo "Starting docker-compose"
	docker-compose -f docker-compose.dev.yml down
