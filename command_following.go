package main

import (
	"context"
	"fmt"

	"github.com/ChernakovEgor/gator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("could not get follows: %v", err)
	}

	fmt.Printf("Listing follows for user %s:\n", s.cfg.User)
	for _, f := range follows {
		fmt.Printf(" - '%s'\n", f.FeedName)
	}
	return nil
}
