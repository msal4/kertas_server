start:
	go run cmd/server/main.go
migrate:
	go run cmd/migrate-schema/main.go
test:
	go test -v ./...
