package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, usr database.User) error {

	ctx := GetContext()

	feeds, errFeeds := s.db.GetFeedFollowsForUser(ctx, usr.ID)

	if errFeeds != nil {
		return errors.New("could not load feeds: " + errFeeds.Error())
	}

	fmt.Printf("All followed feeds for user %s\n", usr.UserName)
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil

}
