package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JMitchell159/gator/internal/config"
	"github.com/JMitchell159/gator/internal/database"
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
	cmds.register("reset", middlewareArgsWarn(handlerReset))
	cmds.register("users", middlewareArgsWarn(handlerUsers))
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", middlewareArgsWarn(handlerFeeds))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareArgsWarn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("agg", handlerAgg)
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

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
