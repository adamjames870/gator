package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	s.db.DeleteAllUsers(context.Background())
	fmt.Println("Deleted all users")
	return nil
}
