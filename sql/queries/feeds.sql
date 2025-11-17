-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, feed_name, feed_url, created_by_user)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedList :many
SELECT feeds.feed_name, feeds.feed_url, created_by_user
FROM feeds
;

-- name: GetFeedByUrl :one

SELECT id, feed_name 
FROM feeds
WHERE feed_url = $1
;
