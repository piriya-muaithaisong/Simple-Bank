package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/piriya-muaithaisong/Simple-Bank/api"
	db "github.com/piriya-muaithaisong/Simple-Bank/db/sqlc"
	"github.com/piriya-muaithaisong/Simple-Bank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
