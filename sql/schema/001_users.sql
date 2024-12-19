-- +goose Up
CREATE TABLE users (
  id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  name text NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;
