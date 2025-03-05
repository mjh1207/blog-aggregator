package main

import (
	"blog-aggregator/internal/config"
	"blog-aggregator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}