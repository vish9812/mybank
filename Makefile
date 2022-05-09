createdb:
	createdb --username=$(USER) --owner=myuser my_bank

dropdb:
	dropdb my_bank

migrateup:
	morph apply up --driver postgres --dsn "postgres://myuser:mypass@localhost:5432/my_bank?sslmode=disable" --path ./db/migration/postgres --number 1

migratedown:
	morph apply down --driver postgres --dsn "postgres://myuser:mypass@localhost:5432/my_bank?sslmode=disable" --path ./db/migration/postgres --number 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc test server