package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/ChernakovEgor/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	freq := cmd.args[0]
	duration, err := time.ParseDuration(freq)
	if err != nil {
		return fmt.Errorf("could not parse duration: %v", err)
	}

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("could not scrape feeds: %v", err)
		}
	}

	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("could not send request: %v", err)
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	var feed RSSFeed
	err = decoder.Decode(&feed)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("could not decode response: %v", err)
	}

	return &feed, nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("could not get next feed to fetch: %v", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("could not fetch feed: %v", err)
	}

	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{Currenttime: time.Now(), FeedID: nextFeed.ID})
	for _, item := range feed.Channel.Items {
		fmt.Println(item.Title)
	}

	return nil
}
