test:
	@go test -v ./...

dev: export DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres
dev:
	go run cmd/grpc-demo-go/main.go