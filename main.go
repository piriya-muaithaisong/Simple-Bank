package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/piriya-muaithaisong/Simple-Bank/api"
	db "github.com/piriya-muaithaisong/Simple-Bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println("cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
