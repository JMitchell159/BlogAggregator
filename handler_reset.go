package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if cmd.args != nil {
		fmt.Println("Warning, reset command takes no arguments:")
		for _, arg := range cmd.args {
			fmt.Printf("%s argument ignored\n", arg)
		}
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("users table reset unsuccessful with error: %v", err)
	}
	fmt.Println("reset users table")

	return nil
}
