package database

import (
	"database/sql"
	"fmt"

	"github.com/abbeyhrt/keep-up-graphql/internal/models"
)

//Store struct for multiple pages
type Store struct {
	db *sql.DB
}

// CreateUser ...
func (s *Store) CreateUser(u models.User) (*models.User, error) {
	sqlCreateUser := `
	INSERT into users
	(name, email, provider)
	VALUES ($1, $2, $3)
	RETURNING id, object_id, name, email, avatar_url, provider, created_at, updated_at
	`

	err := s.db.QueryRow(sqlCreateUser, u.Name, u.Email, u.Provider).Scan(
		&u.ID,
		&u.ObjectID,
		&u.Name,
		&u.Email,
		&u.AvatarURL,
		&u.Provider,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return &u, nil

}
