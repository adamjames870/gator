package main

import (
	"errors"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedList(GetContext())
	if err != nil {
		return errors.New("could not load feds: " + err.Error())
	}

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "URL", "User"})

	for _, row := range feeds {
		tw.AppendRow(table.Row{row.FeedName, row.FeedUrl, row.UserName})
	}
	fmt.Printf("%s\n", tw.Render())
	return nil
}
