package main

import "fmt"

func middlewareArgsWarn(handler func(s *state, cmd command) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if cmd.args != nil {
			fmt.Printf("Warning, %s handler takes no arguments", cmd.name)
			for _, arg := range cmd.args {
				fmt.Printf("%s argument ignored\n", arg)
			}
		}

		return handler(s, cmd)
	}
}
