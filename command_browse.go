package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"log"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)

	if len(cmd.Args) == 1 {
		cmdLimit, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			log.Printf("Expected argument to be an integer. Defaulting to 2 posts")
		} else if cmdLimit <= 0{
			log.Printf("Limit must be positive. Defaulting to 2 posts")
		} else {
			limit = int32(cmdLimit)
		}
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: limit,
	})

	if err != nil {
		return fmt.Errorf("unable to get posts for user %s: %v", user.Name, err)
	}

	for i, post := range posts {
		// feed, err := s.db.GetFeedByUrl(context.Background(), post.Url)
		// if err != nil {
		// 	return fmt.Errorf("unable to get feed by url %s: %v", post.Url, err)
		// }

		title := "No title available"
		if post.Title.Valid {
			title = post.Title.String
		}
		description := "No description available"
		if post.Description.Valid {
			description = post.Description.String
		}
		
		fmt.Printf("\n=== POST %d ===\n", i + 1)
		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Published: %v\n", post.PublishedAt)
		// fmt.Printf("From: %s\n", feed[0].Name)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", description)
	}
	return nil
}