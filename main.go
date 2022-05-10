package main

import (
	"database/sql"
	"log"
	"me/mybank/api"
	db "me/mybank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	driver        = "postgres"
	src           = "postgres://myuser:mypass@localhost:5432/my_bank?sslmode=disable"
	serverAddress = "localhost:7070"
)

func main() {
	conn, err := sql.Open(driver, src)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("server start failed", err)
	}

}
