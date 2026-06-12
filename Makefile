postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres SimpleBankDB

dropdb:
	docker exec -it postgres dropdb SimpleBankDB

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/SimpleBankDB?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/SimpleBankDB?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres creadtedb dropdb migrateup migratedown sqlc test