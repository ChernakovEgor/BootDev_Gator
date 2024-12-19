package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not get feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s - %s - %s\n", feed.Name, feed.Url, feed.Name_2)
	}
	return nil
}
