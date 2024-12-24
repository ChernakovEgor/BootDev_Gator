package main

import (
	"context"
	"fmt"

	"github.com/ChernakovEgor/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	url := cmd.args[0]
	_, err := s.db.DeleteFeedForUser(context.Background(), database.DeleteFeedForUserParams{Name: user.Name, Url: url})
	if err != nil {
		return fmt.Errorf("could not delete feed: %v", err)
	}
	return nil
}
