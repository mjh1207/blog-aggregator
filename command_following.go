package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	// user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	// if err != nil {
	// 	return fmt.Errorf("unable to get data for user %s, error - %v", s.cfg.CurrentUserName, err)
	// }
	
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get follows for user %s, error - %v", s.cfg.CurrentUserName, err)
	}

	fmt.Printf("User %s is currently following:\n", s.cfg.CurrentUserName)
	for _, follow := range follows {
		fmt.Printf("- %s\n", follow.FeedName)
	}
	
	return nil
}