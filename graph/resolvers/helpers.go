package resolvers

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"gitlab.lrz.de/projecthub/gql-api/graph/model"
)

func (r *Resolver) UserFromID(ctx context.Context, id string) (*model.User, error) {
	user, err := r.queries.GetUserByID(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	var name *string
	if user.Name.Valid {
		name = &user.Name.String
	}
	return &model.User{
		ID:       user.ID.String(),
		Username: user.Username,
		Name:     name,
	}, nil
}

func (r *Resolver) UserIdFromusername(ctx context.Context, username string) (string, error) {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
}

func NullstringToPointer(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}
func PointerToNullstring(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}
