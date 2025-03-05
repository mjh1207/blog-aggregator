package main

import (
	"blog-aggregator/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("expected argument user")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		return fmt.Errorf("user %s already exists", cmd.args[0])
	}
	createdUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
	})
	if err != nil { return err }

	// Set current user in the config to current name
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil { return fmt.Errorf("unable to set current user to %s", cmd.args[0])}
	fmt.Println("A user was created successfully")
	fmt.Printf("ID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %v\n", createdUser.ID, createdUser.CreatedAt, createdUser.UpdatedAt, createdUser.Name)
	return nil
}