package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), *s.cfg.Current_User_Name)
	if err != nil {
		return err
	}

	fmt.Printf("%s is following these feeds:\n", *s.cfg.Current_User_Name)
	for _, feed := range follows {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}
