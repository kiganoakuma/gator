-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT 
  inserted_feed_follows.*,
  u.name AS user_name,
  f.name AS feed_name
FROM inserted_feed_follows
INNER JOIN users u ON inserted_feed_follows.user_id = u.id
INNER JOIN feeds f ON inserted_feed_follows.feed_id = f.id;


-- name: GetFeedFollows :many
SELECT * FROM feed_follows;

-- name: GetFeedFollowesForUser :many
SELECT
  f_follows.*,
  feeds.name as feed_name, 
  users.name as user_name 
FROM feed_follows f_follows
INNER JOIN users ON f_follows.user_id = users.id
INNER JOIN feeds ON f_follows.feed_id = feeds.id
WHERE f_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;
