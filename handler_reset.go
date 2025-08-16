package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("users table reset unsuccessful with error: %v", err)
	}
	fmt.Println("reset users table")

	return nil
}
