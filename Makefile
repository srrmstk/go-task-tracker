run:
	@echo "Starting server"
	go run cmd/server/main.go

compose-dev:
	@echo "Starting docker-compose"
	docker-compose -f docker-compose.dev.yml up

compose-dev-d:
	@echo "Starting docker-compose-build"
	docker-compose -f docker-compose.dev.yml up -d

compose-dev-down:
	@echo "Stopping docker-compose"
	docker-compose -f docker-compose.dev.yml down
