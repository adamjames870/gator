package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return errors.New("could not load users: " + err.Error())
	}

	currentUser := s.config.Current_user_name

	for _, usr := range users {
		line := fmt.Sprintf("* %s", usr.UserName)
		if usr.UserName == currentUser {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
}
