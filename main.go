package main

import (
	"blog-aggregator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	// read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	// set current state
	currentState := state {
		cfg: &cfg,
	}

	// register commands
	com := commands {
		make(map[string]func(*state, command) error),
	}
	com.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments provided")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command {
		name: cmdName,
		args: cmdArgs,
	}

	if err := com.run(&currentState, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	
}