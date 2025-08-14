package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JMitchell159/blog_aggregator/internal/config"
	"github.com/JMitchell159/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
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

	db, err := sql.Open("postgres", s.cfg.DB_URL)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	s.db = dbQueries

	cmds := commands{
		handler: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
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
