package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil { return fmt.Errorf("unable to get users list: %v", err)}
	if len(dbUsers) == 0 { 
		fmt.Println("No users currently registered")
		return nil
	}

	for _, user := range dbUsers {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil	
}