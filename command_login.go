package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login handler expects a single argument: username")
	}

	// Check if user exists in the database - error if it doesn't
	dbUser, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil { return fmt.Errorf("user %s does not exist", cmd.args[0])}

	err = s.cfg.SetUser(dbUser.Name)
	if err != nil { return err }

	fmt.Println("User has been set")
	return nil
}