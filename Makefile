start:
	@go run cmd/server/main.go --debug
migrate:
	@go run cmd/migrate-schema/main.go
gen:
	@go generate ./ent
	@go generate .
test:
	@sleep 10
	@go test -v ./...
