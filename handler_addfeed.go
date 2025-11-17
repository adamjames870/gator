package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/jedib0t/go-pretty/v6/table"
)

func handlerAddFeed(s *state, cmd command, usr database.User) error {
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

	newFeed := database.CreateFeedParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		FeedName:      feedName,
		FeedUrl:       feedUrl,
		CreatedByUser: usr.ID,
	}

	fd, errCreateFeed := s.db.CreateFeed(ctx, newFeed)

	if errCreateFeed != nil {
		return errors.New("failed to save feed to DB: " + errCreateFeed.Error())
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"ID", "Created At", "Feed Name", "Feed URL", "Creating User"})
	tw.AppendRow(table.Row{fd.ID, fd.CreatedAt.Format("01-Jan 15:06"), fd.FeedName, fd.FeedUrl, usr.UserName})

	fmt.Printf("%s\n", tw.Render())

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    fd.ID,
	}

	createdFollow, errCreateFollow := s.db.CreateFeedFollow(ctx, newFollow)

	if errCreateFollow != nil {
		return errors.New("failed to create feed_follow: " + errCreateFollow.Error())
	}

	fmt.Printf("Subscribed user %s to feed %s\n", createdFollow.UserName, createdFollow.FeedName)

	return nil

}
