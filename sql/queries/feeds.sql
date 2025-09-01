-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, created_at, updated_at, name, url)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, user_id, created_at, updated_at, name, url;

-- name: DeleteFeeds :exec
TRUNCATE feeds;

-- name: GetFeeds :many
SELECT feeds.name as feed_name, feeds.url, users.name as user_name FROM feeds
INNER JOIN users ON feeds.user_id = users.id;

-- name: GetFeed :one
SELECT feeds.id as feedID, feeds.name as feed_name, feeds.url, users.name as user_name FROM feeds
INNER JOIN users ON feeds.user_id = users.id
WHERE feeds.url = $1;