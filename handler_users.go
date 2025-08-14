package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == *s.cfg.Current_User_Name {
			fmt.Println(" (current)")
			continue
		}
		fmt.Println()
	}

	return nil
}
