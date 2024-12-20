-- +goose Up
CREATE TABLE feed_follows (
  id int NOT NULL GENERATED ALWAYS AS IDENTITY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
