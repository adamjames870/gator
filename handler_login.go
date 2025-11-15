package main

import (
	"context"
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		// Empty or nil arguments, login expects a login name
		return errors.New("no login name provided")
	}

	userName := cmd.args[0]

	usr, check_err := s.db.GetUser(context.Background(), userName)
	if check_err != nil {
		// user does not exist
		return errors.New("user does not exist")
	}

	config.SetUser(*s.config, usr.UserName)
	fmt.Printf("Set username to '%s'\n", usr.UserName)
	return nil
}
