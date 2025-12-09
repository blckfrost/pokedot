.PHONY: redis server client dev build clean

server:
	@echo "Starting Go backend..."
	@go run cmd/server/main.go

client:
	@echo "Starting frontend..."
	@cd client && pnpm install && pnpm dev

redis:
	@echo "Starting redis container..."
	@docker compose up -d

dev:
	@echo "Starting dev environment"
	@make -j3 redis server client
	@go run cmd/server/main.go


build:
	@echo "Building Go backend..."
	@go build -o pokedot cmd/server/main.go

clean:
	@echo "Cleaning..."
	@rm -f pokedot
	@rm -f client/dist
