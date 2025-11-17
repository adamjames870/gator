package main

import (
	"errors"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedList(GetContext())
	if err != nil {
		return errors.New("could not load feeds: " + err.Error())
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "URL", "Creating User"})

	for _, row := range feeds {
		tw.AppendRow(table.Row{row.FeedName, row.FeedUrl, row.CreatedByUser})
	}
	fmt.Printf("%s\n", tw.Render())
	return nil
}
