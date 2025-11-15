package main

import (
	"fmt"
	"github/adamjames870/gator/internal/config"
	"os"
)

func main() {
	config := config.Read()
	state := state{config: &config}
	commands := getNewCommands()
	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Too few arguments entered")
		os.Exit(1)
	}

	cmd := command{name: args[1], args: args[2:]}

	err := commands.run(&state, cmd)
	if err != nil {
		fmt.Printf("Command failed: %s\n", err)
		os.Exit(1)
	}

}
