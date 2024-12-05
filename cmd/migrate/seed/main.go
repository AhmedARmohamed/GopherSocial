package main

import (
	"log"

	"github.com/AhmedARMohamed/social/internal/db"
	"github.com/AhmedARMohamed/social/internal/env"
	"github.com/AhmedARMohamed/social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:admin@localhost:5435/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store)
}
