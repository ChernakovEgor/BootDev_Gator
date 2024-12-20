-- +goose Up
CREATE TABLE feeds (
  id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  name text NOT NULL,
  url text NOT NULL UNIQUE,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
