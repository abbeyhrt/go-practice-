package main

import (
	"log"

	"github.com/abbeyhrt/keep-up-graphql/internal/config"
	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/pubapi/pubapisrv"
)

func main() {
	cfg := config.New()
	store, err := database.New(cfg.Postgres.String())

	if err != nil {
		log.Fatalf("error creating pubapisrv: %v", err)
	}

	srv, err := pubapisrv.New(cfg.Addr, store)

	if err != nil {
		log.Fatalf("error creating databse: %v", err)
	}

	log.Println("Listening at http://" + cfg.Addr)
	log.Fatal(srv.ListenAndServe())
}
