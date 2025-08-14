package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/JMitchell159/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("the register handler expects a single argument, the username")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		return fmt.Errorf("username %s already exists in the database", cmd.args[0])
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User was created with the following data:\nID: %v\nCreated: %v\nUpdated: %v\nName: %s\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	return nil
}
