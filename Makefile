db_migrate:
	go run cmd/db/main.go -env=test init
	go run cmd/db/main.go -env=test migrate

test:
	go test ./tests