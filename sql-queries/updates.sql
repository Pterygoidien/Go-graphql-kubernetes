-- name: GetUpdatesForProject :many
SELECT * FROM updates
WHERE project_id = $1 AND timestamp <= $2
ORDER BY timestamp DESC
LIMIT $3;
-- name: CreateUpdateForProject :exec
INSERT INTO updates (project_id, content, images)
VALUES($1, $2, $3);
-- name: DeleteUpdate :exec
DELETE FROM updates
WHERE id = $1;
-- name: DeleteTwitterUpdatesForProject :exec
DELETE from updates
WHERE project_id=$1 AND source='twitter';