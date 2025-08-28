package main

import (
	"log"

	"github.com/AndreasNugoho/go-social/internal/db"
	"github.com/AndreasNugoho/go-social/internal/env"
	"github.com/AndreasNugoho/go-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost:5438/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
