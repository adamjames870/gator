package rss

import (
	"context"
	"encoding/xml"
	"errors"
	"html"
	"io"
	"net/http"
	"time"
)

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	request, reqErr := http.NewRequestWithContext(ctx, "", feedUrl, nil)
	if reqErr != nil {
		return nil, errors.New("error opening request: " + reqErr.Error())
	}

	request.Header.Set("User-Agent", "gator")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	response, respErr := client.Do(request)
	if respErr != nil {
		return nil, errors.New("error fetching feed: " + respErr.Error())
	}

	body, _ := io.ReadAll(response.Body)
	var responseData RSSFeed

	unmErr := xml.Unmarshal(body, &responseData)

	if unmErr != nil {
		return nil, errors.New("error unmarshalling data: " + unmErr.Error())
	}

	cleanText(&responseData.Channel.Title)
	cleanText(&responseData.Channel.Description)

	for _, item := range responseData.Channel.Item {
		cleanText(&item.Title)
		cleanText(&item.Description)
	}

	return &responseData, nil

}

func cleanText(txt *string) {
	*txt = html.UnescapeString(*txt)
}
