package main

import (
	"github/adamjames870/gator/internal/config"
	"github/adamjames870/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
