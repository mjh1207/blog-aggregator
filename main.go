package main

import (
	"blog-aggregator/internal/config"
	"blog-aggregator/internal/database"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("Error: unable to open database connection", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	
	// set current state
	currentState := state {
		db: dbQueries,
		cfg: &cfg,
	}

	// register commands
	com := commands {
		handlers: make(map[string]func(*state, command) error),
	}
	com.register("login", handlerLogin)
	com.register("register", handlerRegister)
	com.register("reset", handlerReset)

	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments provided")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command {
		Name: cmdName,
		Args: cmdArgs,
	}

	if err := com.run(&currentState, cmd); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	os.Exit(0)
	
}