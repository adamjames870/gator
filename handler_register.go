package main

import (
	"context"
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/config"
	"github/adamjames870/gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		// Empty or nil arguments, login expects a login name
		return errors.New("no name provided")
	}

	userName := cmd.args[0]

	_, check_err := s.db.GetUser(context.Background(), userName)
	if check_err == nil {
		// user exists, do not add
		return errors.New("user already exists")
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserName:  userName,
	}

	usr, create_err := s.db.CreateUser(context.Background(), newUser)
	if create_err != nil {
		return errors.New("failed to create user: " + create_err.Error())
	}

	fmt.Println(usr)
	config.SetUser(*s.config, usr.UserName)
	return nil

}
