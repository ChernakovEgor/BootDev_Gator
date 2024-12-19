package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAgg(s *state, _ command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)

	fmt.Printf("feed: %v\n", feed)

	return err
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
