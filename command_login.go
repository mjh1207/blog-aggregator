package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login handler expects a single argument: username")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil { return err }

	fmt.Println("User has been set")
	return nil
}