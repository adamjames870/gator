package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
	"github/adamjames870/gator/internal/rss"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func handlerScrapeFeeds(s *state, cmd command) error {
	ctx := GetContext()
	feedToFetch, errUrl := s.db.GetNextFeedToFecth(ctx)
	if errUrl != nil {
		return errors.New("failed to select URL: " + errUrl.Error())
	}
	data, errData := rss.FetchFeed(ctx, feedToFetch.FeedUrl)
	if errData != nil {
		return errors.New("failed to load feed: " + errData.Error())
	}

	feedUpdate := database.MarkFeedFetchedParams{
		ID:        feedToFetch.ID,
		UpdatedAt: time.Now(),
	}
	updatedFetch, errUpdatedFetch := s.db.MarkFeedFetched(ctx, feedUpdate)
	if errUpdatedFetch != nil {
		return errors.New("failed to mark feed updated: " + errUpdatedFetch.Error())
	}

	fmt.Printf("Items from %s(%s) | Marked updated at %s\n", data.Channel.Title, data.Channel.Description, updatedFetch.LastFetchedAt.Time.Format("15:04"))
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Title"})
	for _, item := range data.Channel.Item {
		tw.AppendRow(table.Row{item.Title})
	}
	fmt.Printf("%s\n", tw.Render())
	return nil
}
