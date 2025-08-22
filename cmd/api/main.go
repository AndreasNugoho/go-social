package main

import (
	"log"

	"github.com/AndreasNugoho/go-social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":9090"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
