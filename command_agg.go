package main

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected argument: duration string")
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("expected argument: duration string")
	}

	fmt.Printf("Collecting feeds every %v\n", interval)

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("there was an error scraping the next feed")
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now().UTC(),
		LastFetchedAt: sql.NullTime{time.Now().UTC(), true},
		ID: nextFeed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Fetching feed from %s\n", nextFeed.Url)
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed from %s - Error: %v", nextFeed.Url, err)
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("- %s\n", item.Title)
	}

	return nil
}