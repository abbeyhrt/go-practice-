package pubapisrv

import (
	"net/http"
	"time"

	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/handlers"
)

// New server
func New(addr string, store database.DAL) (*http.Server, error) {
	handler := handlers.New(store)
	srv := http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	return &srv, nil
}
