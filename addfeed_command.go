package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ChernakovEgor/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("not enough arguments")
	}
	title := cmd.args[0]
	url := cmd.args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.User)
	if err != nil {
		return fmt.Errorf("could not get user: %v", err)
	}

	feedParams := database.AddFeedParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: title, Url: url, UserID: user.ID}
	_, err = s.db.AddFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("could not add feed: %v", err)
	}

	return nil
}
