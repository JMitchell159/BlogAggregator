package main

import (
	"fmt"
	"os"

	"github.com/JMitchell159/blog_aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	s := state{
		cfg: c,
	}

	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	inputs := os.Args
	if len(inputs) < 2 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	cmd := command{
		name: inputs[1],
		args: nil,
	}
	if len(inputs) > 2 {
		cmd.args = inputs[2:]
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
