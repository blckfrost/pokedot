.PHONY: compose server client dev build clean

server:
	@go mod tidy
	@echo "Starting Go backend..."
	@go run cmd/server/main.go

client:
	@echo "Starting frontend..."
	@cd client && pnpm install && pnpm dev

compose:
	@echo "Starting redis+psql containers..."
	@docker compose up -d

dev:
	@echo "Starting dev environment"
	@make -j3 compose server client
	@go run cmd/server/main.go


build:
	@echo "Building Go backend..."
	@go build -o pokedot cmd/server/main.go

clean:
	@echo "Cleaning..."
	@rm -f pokedot
	@rm -f client/dist
