package main

import (
	"blog-aggregator/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Println(cfg)
	err = cfg.SetUser("Mitch")
	if err != nil {
		log.Fatalf("err setting username: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Println(cfg)
	
}