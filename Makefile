start:
	@go run cmd/server/main.go --debug
migrate:
	@go run cmd/migrate-schema/main.go
gen:
	@go generate ./ent
	@go generate .
seed:
	@go run cmd/seed/main.go
test:
	@go test -cover -v ./...
