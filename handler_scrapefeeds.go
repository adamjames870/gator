package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
	"github/adamjames870/gator/internal/rss"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
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
	_, errUpdatedFetch := s.db.MarkFeedFetched(ctx, feedUpdate)
	if errUpdatedFetch != nil {
		return errors.New("failed to mark feed updated: " + errUpdatedFetch.Error())
	}

	for _, item := range data.Channel.Item {

		dt, errDt := dateparse.ParseAny(item.PubDate)

		if errDt != nil {
			fmt.Printf("Could not parse datetime %s with error: %s. Post will not be saved.\n", item.PubDate, errDt.Error())
			continue
		}

		new_post := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: dt,
			FeedID:      feedUpdate.ID,
		}

		post, errPost := s.db.CreatePost(ctx, new_post)
		if errPost != nil {
			if strings.Contains(errPost.Error(), "violates unique constraint \"posts_url_key\"") {
				fmt.Printf("Not saving post (already saved in database)\nPost: %s\nFrom feed: %s\n", new_post.Title, data.Channel.Title)
			} else {
				fmt.Printf("Failed to save post: %s\n", errPost.Error())
			}
		} else {
			fmt.Printf("Saved post to database\nPost: %s\nFrom feed: %s\n", post.Title, data.Channel.Title)
		}
		fmt.Println("--------------------------------------------------------")

	}

	return nil
}
