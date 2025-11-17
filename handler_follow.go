package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, usr database.User) error {

	if len(cmd.args) == 0 {
		// Empty or nil arguments, follow expects a url to follow
		return errors.New("no url provided")
	}

	ctx := GetContext()

	feedUrl := cmd.args[0]
	feedIdToFollow, errGetFeed := s.db.GetFeedByUrl(ctx, feedUrl)

	if errGetFeed != nil {
		return errors.New("unable to get feed from url: " + errGetFeed.Error())
	}

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    feedIdToFollow.ID,
	}

	createdFollow, errCreateFollow := s.db.CreateFeedFollow(ctx, newFollow)

	if errCreateFollow != nil {
		return errors.New("failed to create feed_follow: " + errCreateFollow.Error())
	}

	fmt.Printf("Subscribed user %s to feed %s\n", createdFollow.UserName, createdFollow.FeedName)
	return nil

}
