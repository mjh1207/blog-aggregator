package main

import (
	"blog-aggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
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
			fmt.Printf("%v", err)
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
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC(), Valid: true},
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
		fmt.Println(item.Link)
		pubAt, err := dateparse.ParseAny(item.PubDate)

		if err != nil {
			log.Printf("Failed to parse date %s: %v", item.PubDate, err)
			pubAt = time.Now().UTC()
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: sql.NullString{
				String: strings.TrimSpace(item.Title),
				Valid: item.Title != "",
			},
			Url: item.Link,
			Description: sql.NullString{
				String: strings.TrimSpace(item.Description),
				Valid: item.Description != "",
			},
			PublishedAt: sql.NullTime{
				Time: pubAt,
				Valid: !pubAt.IsZero(),
			},
			FeedID: nextFeed.ID,
		})
		if err != nil && strings.Contains(err.Error(), "duplicate key") {
			continue
		} else if err != nil {
			log.Printf("Failed to create post: %v", err)
		}
	}

	return nil
}