package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	errDelete := s.db.DeleteAllUsers(context.Background())
	if errDelete != nil {
		return errors.New("failed to delete users: " + errDelete.Error())
	}
	fmt.Println("Deleted all users")
	return nil
}
