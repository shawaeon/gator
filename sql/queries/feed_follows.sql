-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
) 
SELECT inserted.*, users.name AS user_name, feeds.name AS feed_name
FROM inserted
LEFT JOIN users ON feed_follows.user_id = users.id
LEFT JOIN feeds ON feed_follows.feed_id = feeds.id
;