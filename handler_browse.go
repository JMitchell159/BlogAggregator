package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/JMitchell159/blog_aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	lim := 2
	if len(cmd.args) >= 1 {
		temp, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		lim = temp
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(lim),
	})
	if err != nil {
		return err
	}

	for i, post := range posts {
		fmt.Printf("Post %d:\n", i+1)
		fmt.Printf("Title: %s\n", post.Title.String)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Publish Date: %v\n", post.PublishedAt.Format(time.RFC1123Z))
	}

	return nil
}
