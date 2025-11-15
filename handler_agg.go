package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	data, err := rss.FetchFeed(GetContext(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return errors.New("failed to get data: " + err.Error())
	}
	fmt.Println(data)
	return nil
}
