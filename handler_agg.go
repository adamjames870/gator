package main

import (
	"errors"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		// Empty or nil arguments, expects a timer
		return errors.New("no timer period provided")
	}

	period, errPeriod := time.ParseDuration(cmd.args[0])
	if errPeriod != nil {
		return errors.New("failed to parse duration: " + errPeriod.Error())
	}

	fmt.Printf("Collecting feeds every %s\n", period)

	ticker := time.NewTicker(period)
	for ; ; <-ticker.C {
		handlerScrapeFeeds(s, cmd)
	}

}
