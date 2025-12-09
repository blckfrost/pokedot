run-be:
	@go run cmd/server/main.go

run:
	

redis:
	@docker compose up -d

clean:
	@echo "Cleaning..."
	@rm -f main
	@rm -f client/dist
