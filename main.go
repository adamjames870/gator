package main

import (
	"errors"
	"fmt"
)

func main() {
	config := ReadConfig()
	fmt.Println(config.Db_url)
	config.Current_user_name = "Adam"
	WriteConfig(config)
	newConfig := ReadConfig()
	fmt.Println(newConfig.Current_user_name)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		// Empty or nil arguments, login expects a login name
		return errors.New("no login name provided")
	}
	s.config.Current_user_name = cmd.args[0]
	fmt.Printf("Set username to '%s'\n", s.config.Current_user_name)
	return nil
}
