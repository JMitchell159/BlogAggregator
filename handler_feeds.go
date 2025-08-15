package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if cmd.args != nil {
		fmt.Println("Warning, feeds command takes no arguments:")
		for _, arg := range cmd.args {
			fmt.Printf("%s argument ignored\n", arg)
		}
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("===========================")
	for _, feed := range feeds {
		user, err := s.db.GetUserFromID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed Name: %s\n", feed.Name)
		fmt.Printf("Feed URL: %s\n", feed.Url)
		fmt.Printf("Feed Creator: %s\n", user.Name)
		fmt.Println("===========================")
	}

	return nil
}
