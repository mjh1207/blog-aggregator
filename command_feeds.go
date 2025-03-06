package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("expected 0 arguments but got: %v", len(cmd.Args))
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get feeds - error: %v", err)
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	for _, feed := range feeds {
		name, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil { return fmt.Errorf("unable to get user for feed - Error: %v", err)}
		fmt.Println("Name:", feed.Name)
		fmt.Println("URL:", feed.Url)
		fmt.Println("User name:", name)
	}

	return nil
}