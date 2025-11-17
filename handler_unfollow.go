package main

import (
	"errors"
	"github/adamjames870/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, usr database.User) error {
	if len(cmd.args) == 0 {
		// Empty or nil arguments, follow expects a url to unfollow
		return errors.New("no url provided")
	}

	ctx := GetContext()

	feedUrl := cmd.args[0]
	feedIdToUnFollow, errGetFeed := s.db.GetFeedByUrl(ctx, feedUrl)

	if errGetFeed != nil {
		return errors.New("unable to get feed from url: " + errGetFeed.Error())
	}

	prms := database.DeleteFeedFollowParams{
		UserID: usr.ID,
		FeedID: feedIdToUnFollow.ID,
	}

	unfollowErr := s.db.DeleteFeedFollow(ctx, prms)
	if unfollowErr != nil {
		return errors.New("unable to unfollow: " + unfollowErr.Error())
	}

	return nil

}
