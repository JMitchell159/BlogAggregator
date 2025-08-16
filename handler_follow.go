package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JMitchell159/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if cmd.args == nil {
		return errors.New("the follow handler expects a single argument, the feed url")
	}

	if len(cmd.args) > 1 {
		fmt.Printf("Warning, all arguments except for %s will be ignored\n", cmd.args[0])
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed Follow created with the following data:")
	fmt.Printf("Feed Name: %s\n", feed_follow.FeedName)
	fmt.Printf("Follower: %s\n", feed_follow.UserName)

	return nil
}
