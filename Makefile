run:
	go run ./cmd/app/main.go

lint:
	golangci-lint run ./...

ent-gen:
	go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/lock,sql/upsert ./internal/ent/schema
