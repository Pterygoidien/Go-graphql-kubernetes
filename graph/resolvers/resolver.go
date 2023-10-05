package resolvers

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gitlab.lrz.de/projecthub/gql-api/sqlc"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	queries *sqlc.Queries
}

func NewResolver(connstring string) (*Resolver, error) {
	//db, err := pgxpool.Connect(context.Background(), connstring)
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		return nil, err
	}
	queries := sqlc.New(db)
	return &Resolver{
		queries: queries,
	}, nil
}

func NewDemoResolver() *Resolver {
	return &Resolver{}
}
