package main

import (
	"log"

	"github.com/AndreasNugoho/go-social/internal/env"
	"github.com/AndreasNugoho/go-social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":9090"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
