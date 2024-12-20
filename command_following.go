package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, _ command) error {

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.User)
	if err != nil {
		return fmt.Errorf("could not get follows: %v", err)
	}

	fmt.Printf("Listing follows for user %s:\n", s.cfg.User)
	for _, f := range follows {
		fmt.Printf(" - '%s'\n", f.FeedName)
	}
	return nil
}
