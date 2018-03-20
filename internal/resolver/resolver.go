package resolver

import "github.com/abbeyhrt/keep-up-graphql/internal/database"

//New store created for resolvers
func New(store database.DAL) *Resolver {
	return &Resolver{store}
}

//Resolver struct created for store
type Resolver struct {
	db database.DAL
}
