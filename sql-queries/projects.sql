-- name: GetProjects :many
SELECT *
FROM projects
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
-- name: CreateProject :one
INSERT into projects (name, description, creator)
VALUES ($1, $2, $3)
RETURNING id;
-- name: GetProjectByID :one
SELECT *
FROM projects
WHERE id = $1;
-- name: GetProjectsByUserID :many
SELECT *
FROM projects
WHERE creator = $1;
-- name: GetParticipantsOfProject :many
SELECT DISTINCT users.*
FROM users
    INNER JOIN participations ON users.id = participations.user_id
WHERE participations.project_id = $1;
-- name: AddParticipantToProject :exec
INSERT INTO participations (user_id, project_id)
VALUES($1, $2);
-- name: RemoveParticipantFromProject :exec
DELETE FROM participations
WHERE user_id = $1
    AND project_id = $2;
-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1;
-- name: UpdateProjectName :exec
UPDATE projects
set name = $2
WHERE id = $1;
-- name: UpdateProjectDescription :exec
UPDATE projects
set description = $2
WHERE id = $1;
-- name: AddImageToProject :exec
INSERT INTO images (project, url, description, priority)
VALUES($1, $2, $3, $4);
-- name: DeleteImage :exec
DELETE FROM images
WHERE id = $1;
-- name: GetImagesOfProject :many
SELECT *
FROM images
WHERE project = $1
ORDER BY priority DESC;
-- name: UpdateImagePriority :exec
UPDATE images
set priority = $2
WHERE id = $1;
-- name: UpdateImageDescription :exec
UPDATE images
set description = $2
WHERE id = $1;
-- name: SearchProjects :many
SELECT *
FROM projects_ranked
WHERE name LIKE $3
    OR description LIKE $3
ORDER BY score DESC
LIMIT $1 OFFSET $2;
-- name: TagProject :exec
INSERT INTO taggings (name, project_id)
VALUES($1, $2);
-- name: UntagProject :exec
DELETE FROM taggings
WHERE name = $1
    AND project_id = $2;
-- name: SearchProjectsWithTag :many
SELECT p.*
FROM projects_ranked p
    INNER JOIN taggings ON p.id = taggings.project_id
WHERE taggings.name = $3
    AND (
        p.name LIKE $4
        OR p.description LIKE $4
    )
ORDER BY score DESC
LIMIT $1 OFFSET $2;
-- name: GetProjectTags :many
SELECT name
from taggings
WHERE project_id = $1;
-- name: GetTagsByCount :many
SELECT name,
    COUNT(*) as count
FROM taggings
GROUP BY name
ORDER BY count DESC
LIMIT $1 OFFSET $2;
-- name: SetProjectLocation :exec
UPDATE projects
set location = $2
WHERE id = $1;
-- name: SetTwitterDatasource :exec
UPDATE projects
set twitter_account = $2
WHERE id = $1;
-- name: SetProjectVisibility :exec
UPDATE projects
set public = $2
WHERE id = $1;