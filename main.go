package main

import (
	"database/sql"
	"fmt"
	"github/adamjames870/gator/internal/config"
	"github/adamjames870/gator/internal/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	config := config.Read()
	state := state{config: &config}
	commands := getNewCommands()
	commands = commands.registerCommands()

	dbUrl := state.config.Db_url
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Printf("Error loading DB: %s\n", err)
	}
	dbQueries := database.New(db)
	state.db = dbQueries

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Too few arguments entered")
		os.Exit(1)
	}

	cmd := command{name: args[1], args: args[2:]}

	err = commands.run(&state, cmd)
	if err != nil {
		fmt.Printf("Command failed: %s\n", err)
		os.Exit(1)
	}

}
