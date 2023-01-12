postgres_run:
	sudo docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:15-alpine

postgres_start:
	sudo docker start 97

createdb:
	sudo docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path ./db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mock_db -destination ./db/mock/store.go github.com/piriya-muaithaisong/Simple-Bank/db/sqlc Store

.PHONY: postgres_run createdb dropdb migrateup migratedown sqlc db server postgres_start mock