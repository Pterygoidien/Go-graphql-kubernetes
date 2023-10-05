package model

import (
	"time"

	"gitlab.lrz.de/projecthub/gql-api/sqlc"
)

type Project struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Languages   []string   `json:"languages"`
	Location    *Location  `json:"location"`
	CreatedAt   *time.Time `json:"createdAt"`
	LastUpdated *time.Time `json:"lastUpdated"`
	Public      bool       `json:"public"`
	CreatorID   string
}

func ProjectFromDBProject(dbProject sqlc.Project) *Project {
	createdAt := &dbProject.CreatedAt.Time
	if !dbProject.CreatedAt.Valid {
		createdAt = nil
	}
	var location *Location
	if dbProject.Location.Valid {
		location = &Location{
			Name: dbProject.Location.String,
		}
	}
	return &Project{
		ID:          dbProject.ID.String(),
		Name:        dbProject.Name,
		Description: dbProject.Description,
		CreatorID:   dbProject.Creator.String(),
		CreatedAt:   createdAt,
		Languages:   []string{},
		Location:    location,
		Public:      dbProject.Public,
	}
}

func ProjectFromRankedDBProject(dbProject sqlc.ProjectsRanked) *Project {
	createdAt := &dbProject.CreatedAt.Time
	if !dbProject.CreatedAt.Valid {
		createdAt = nil
	}
	lastUpdated := &dbProject.LastChange.Time
	if !dbProject.LastChange.Valid {
		lastUpdated = nil
	}
	var location *Location
	if dbProject.Location.Valid {
		location = &Location{
			Name: dbProject.Location.String,
		}
	}
	return &Project{
		ID:          dbProject.ID.String(),
		Name:        dbProject.Name,
		Description: dbProject.Description,
		CreatorID:   dbProject.Creator.String(),
		CreatedAt:   createdAt,
		Languages:   []string{},
		Location:    location,
		LastUpdated: lastUpdated,
		Public: 	dbProject.Public,
	}
}

func ProjectsFromRankedDBProjects(dbProject []sqlc.ProjectsRanked) []*Project {
	projects := make([]*Project, len(dbProject))
	for i, dbProject := range dbProject {
		projects[i] = ProjectFromRankedDBProject(dbProject)
	}
	return projects
}

func ProjectsFromDBProjects(dbProject []sqlc.Project) []*Project {
	projects := make([]*Project, len(dbProject))
	for i, dbProject := range dbProject {
		projects[i] = ProjectFromDBProject(dbProject)
	}
	return projects
}

type ProjectMutation struct {
	ID string
}
