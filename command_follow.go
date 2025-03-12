package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow( s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected argument: url")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	if len(feed) != 1 {
		return fmt.Errorf("feed does not exist for url: %s", cmd.Args[0])
	}
	

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: feed[0].ID,
		UserID: user.ID,

	})
	if err != nil {
		return fmt.Errorf("unable to add follow - %v", err)
	}

	fmt.Printf("User %s is now following feed %s", user.Name, feed[0].Name)

	return nil
}
