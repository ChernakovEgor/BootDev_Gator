// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
with feed_insert as (
  INSERT INTO feed_follows (created_at, updated_at, user_id, feed_id)
   SELECT $1 as created_at,
          $2 as updated_at,
          u.id as user_id,
          f.id as feed_id
    FROM users u
    JOIN feeds f ON f.url = $3
   WHERE u.name = $4
RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT fi.id, fi.created_at, fi.updated_at, fi.user_id, fi.feed_id,
       u.name as user_name,
       f.name as feed_name
  FROM feed_insert fi
  JOIN users u ON fi.user_id = u.id
  JOIN feeds f ON fi.feed_id = f.id
`

type CreateFeedFollowParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Url       string
	Name      string
}

type CreateFeedFollowRow struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	UserName  string
	FeedName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Url,
		arg.Name,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.UserName,
		&i.FeedName,
	)
	return i, err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT u.name as user_name,
       f.name as feed_name
  FROM users u 
  JOIN feed_follows ff ON ff.user_id = u.id
  JOIN feeds f ON f.id = ff.feed_id
 WHERE u.name = $1
`

type GetFeedFollowsForUserRow struct {
	UserName string
	FeedName string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, name string) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(&i.UserName, &i.FeedName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
