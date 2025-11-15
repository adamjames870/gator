package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
	"net/url"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		// two arguments required - name and url
		return errors.New("need two args - name and feed")
	}

	feedName := cmd.args[0]
	feedUrl := cmd.args[1]

	_, checkUrl := url.ParseRequestURI(feedUrl)
	if checkUrl != nil {
		return errors.New("badly formed URL")
	}

	ctx := GetContext()

	currentUser := s.config.Current_user_name
	currentUserId, idErr := s.db.GetUserIdFromName(ctx, currentUser)

	if idErr != nil {
		return errors.New("failed to load user id")
	}

	newFeed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedName:  feedName,
		FeedUrl:   feedUrl,
		UserID:    currentUserId,
	}

	fd, fdErr := s.db.CreateFeed(ctx, newFeed)

	if fdErr != nil {
		return errors.New("failed to save feed to DB")
	}

	fmt.Println(fd)
	return nil

}
