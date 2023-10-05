package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gitlab.lrz.de/projecthub/gql-api/graph/model"
	"gitlab.lrz.de/projecthub/gql-api/sqlc"
)

func (r *projectResolver) Updates(ctx context.Context, obj *model.Project, until *time.Time, count int) ([]*model.Update, error) {
	if until == nil {
		tmp := time.Now()
		until = &tmp
	}
	dbUpdates, err := r.queries.GetUpdatesForProject(context.Background(), sqlc.GetUpdatesForProjectParams{
		ProjectID: uuid.MustParse(obj.ID),
		Timestamp: *until,
		Limit:     int32(count),
	})
	if err != nil {
		return nil, err
	}
	updates := make([]*model.Update, len(dbUpdates))
	for i, dbUpdate := range dbUpdates {
		updates[i] = &model.Update{
			ID:      dbUpdate.ID.String(),
			Content: dbUpdate.Content,
			Origin:  dbUpdate.Source,
			Images:  strings.Split(dbUpdate.Images, ","),
			Time:    dbUpdate.Timestamp,
		}
	}
	return updates, nil
}

func (r *projectResolver) TwitterUpdateSource(ctx context.Context, obj *model.Project) (*string, error) {
	dbProject, err := r.queries.GetProjectByID(context.Background(), uuid.MustParse(obj.ID))
	if err != nil {
		return nil, err
	}
	if !dbProject.TwitterAccount.Valid {
		return nil, nil
	}
	return &dbProject.TwitterAccount.String, nil
}

func (r *projectMutationResolver) CreateUpdate(ctx context.Context, obj *model.ProjectMutation, content string, images []string) (bool, error) {
	err := r.queries.CreateUpdateForProject(context.Background(), sqlc.CreateUpdateForProjectParams{
		ProjectID: uuid.MustParse(obj.ID),
		Content:   content,
		Images:    strings.Join(images, ","),
	})
	return err == nil, err
}

func (r *projectMutationResolver) DeleteUpdate(ctx context.Context, obj *model.ProjectMutation, id string) (bool, error) {
	err := r.queries.DeleteUpdate(context.Background(), uuid.MustParse(obj.ID))
	return err == nil, err
}

func (r *projectMutationResolver) SetTwitterUpdateSource(ctx context.Context, obj *model.ProjectMutation, username *string) (bool, error) {
	err := r.queries.DeleteTwitterUpdatesForProject(context.Background(), uuid.MustParse(obj.ID))
	if err != nil {
		return false, fmt.Errorf("could not delete old tweets")
	}
	err = r.queries.SetTwitterDatasource(context.Background(), sqlc.SetTwitterDatasourceParams{
		ID:             uuid.MustParse(obj.ID),
		TwitterAccount: PointerToNullstring(username),
	})
	return err == nil, err
}
