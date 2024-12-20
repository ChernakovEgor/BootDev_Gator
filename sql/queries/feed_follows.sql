-- name: CreateFeedFollow :one
with feed_insert as (
  INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
   SELECT $1 as created_at,
          $2 as updated_at,
          u.id as user_id,
          f.id as feed_id
    FROM users u
    JOIN feeds f ON f.url = $3
   WHERE u.name = $4
RETURNING *
)
SELECT fi.*,
       u.name as user_name,
       f.name as feed_name
  FROM feed_insert fi
  JOIN users u ON fi.user_id = u.id
  JOIN feeds f ON fi.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT u.name as user_name,
       f.name as feed_name
  FROM users u 
  JOIN feed_follows ff ON ff.user_id = u.id
  JOIN feeds f ON f.id = ff.feed_id
 WHERE u.name = $1;
