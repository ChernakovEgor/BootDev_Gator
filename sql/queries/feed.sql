-- name: AddFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING *;

-- name: GetFeeds :many
SELECT f.name,
       f.url,
       u.name
  FROM feeds f
JOIN users u ON u.id = f.user_id;

-- name: MarkFeedFetched :exec
UPDATE feeds
   SET updated_at = @currentTime,
       last_fetched_at = @currentTime
 WHERE feeds.id = @feed_id;

-- name: GetNextFeedToFetch :one
SELECT *
  FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;
