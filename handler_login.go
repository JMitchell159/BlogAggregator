package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("the login handler expects a single argument, the username")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Current User: %s\n", cmd.args[0])
	return nil
}
