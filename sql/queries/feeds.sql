-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name AS username
FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT id, name
FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT feeds.id, feeds.url
FROM feeds
INNER JOIN users ON feeds.user_id = users.id
WHERE users.name = $1
ORDER BY feeds.last_fetched_at ASC NULLS FIRST, feeds.created_at ASC
LIMIT 1;
