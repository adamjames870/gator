package main

import (
	"errors"
	"github/adamjames870/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	ctx := GetContext()

	return func(s *state, cmd command) error {
		currentUserName := s.config.Current_user_name
		currentUserId, errGetUser := s.db.GetUserIdFromName(ctx, currentUserName)
		if errGetUser != nil {
			return errors.New("failed to load user id" + errGetUser.Error())
		}
		user := database.User{ID: currentUserId, UserName: currentUserName}
		return handler(s, cmd, user)
	}

}
