package main

import (
	"context"
	"errors"

	"github.com/JMitchell159/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if cmd.args == nil {
		return errors.New("the unfollow handler expects a single argument, the feed's url")
	}

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	})

	return err
}
