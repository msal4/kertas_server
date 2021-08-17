start:
	@go run cmd/server/main.go --debug
migrate:
	@go run cmd/migrate-schema/main.go
gen:
	@go generate ./ent
	@go generate .
test:
	@go test -v ./...
