package main

import (
	"errors"
	"fmt"
	"github/adamjames870/gator/internal/database"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func handlerBrowse(s *state, cmd command, usr database.User) error {

	ctx := GetContext()

	postLimit := 2
	postPageNum := 1

	if len(cmd.args) > 0 {
		// user has provided a number of posts to view
		lim, errLim := strconv.Atoi(cmd.args[0])
		if errLim != nil {
			return errors.New("provide a numeric limit: " + errLim.Error())
		}
		postLimit = lim
	}

	if len(cmd.args) > 1 {
		// user has also provided a page numer
		pg, errPg := strconv.Atoi(cmd.args[1])
		if errPg != nil {
			return errors.New("provide a numeric page number: " + errPg.Error())
		}
		postPageNum = pg
	}

	ofst := (postPageNum - 1) * postLimit

	postsParams := database.GetPostsForUserParams{
		UserID: usr.ID,
		Limit:  int32(postLimit),
		Offset: int32(ofst),
	}

	posts, errPosts := s.db.GetPostsForUser(ctx, postsParams)

	if errPosts != nil {
		return errors.New("Failed to load posts: " + errPosts.Error())
	}

	for _, post := range posts {
		desc, _ := stripHTML(post.Description.String)
		fmt.Printf("\033[1m%s\033[0m\n%s\n------------------------------------------------\n", post.Title, desc)
	}
	return nil

}

func stripHTML(input string) (string, error) {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", err
	}

	var output strings.Builder
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.TextNode {
			// Append text content
			output.WriteString(n.Data)
		}
		// Recursively process child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return strings.TrimSpace(output.String()), nil
}
