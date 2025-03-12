package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("expected arguments: url")
	}

	err := s.db.Unfollow(context.Background(), database.UnfollowParams{
		Name: user.Name,
		Url: cmd.Args[0],
	})

	if err != nil {
		return fmt.Errorf("unable to unfollow user %s from feed %s - Error %v", user.Name, cmd.Args[0], err)
	}

	return nil
}