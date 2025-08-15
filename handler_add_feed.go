package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JMitchell159/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("the addfeed handler expects 2 arguments, the name of the feed, and the url of the feed")
	}
	feedUser, err := s.db.GetUser(context.Background(), *s.cfg.Current_User_Name)
	if err != nil {
		return err
	}
	if len(cmd.args) > 2 {
		fmt.Printf("Warning, only %v & %v will be used as arguments, all others will be ignored", cmd.args[0], cmd.args[1])
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    feedUser.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed was created for %s with the following data:\n", feedUser.Name)
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("Created: %v\n", feed.CreatedAt)
	fmt.Printf("Updated: %v\n", feed.UpdatedAt)
	fmt.Printf("Name: %v\n", feed.Name)
	fmt.Printf("URL: %v\n", feed.Url)
	fmt.Printf("User ID: %v\n", feed.UserID)
	fmt.Printf("User: %v\n", feedUser.Name)

	return nil
}
