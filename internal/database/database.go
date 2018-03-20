package database

import (
	"database/sql"

	"fmt"

	"github.com/abbeyhrt/keep-up-graphql/internal/models"
	//i need it
	_ "github.com/lib/pq"
)

//DAL interface for Db
type DAL interface {
	CreateUser(user models.User) (*models.User, error)
}

//New connection to DB
func New(connStr string) (DAL, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening db connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %v", err)
	}

	return &Store{db}, nil
}
