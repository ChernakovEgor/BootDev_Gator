package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ChernakovEgor/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("could not follow: not enough arguments")
	}

	currentUser := s.cfg.User
	url := cmd.args[0]
	followParams := database.CreateFeedFollowParams{CreatedAt: time.Now(), UpdatedAt: time.Now(), Url: url, Name: currentUser}
	res, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("could not follow feed: %v", err)
	}

	fmt.Printf("User '%s' started following '%s'\n", res.UserName, res.FeedName)

	return nil
}
