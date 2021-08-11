start:
	go run cmd/server/main.go
migrate:
	go run cmd/migrate-schema/main.go
gen:
	go generate .
test:
	go test -v ./...
