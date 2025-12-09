run:
	@go run cmd/server/main.go

redis:
	@docker compose up -d

clean:
	@echo "Cleaning..."
	@rm -f main
	@rm -f client/dist
