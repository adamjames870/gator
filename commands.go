package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmd_names map[string]func(*state, command) error
}

func (cmds commands) registerCommands() commands {
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("scrape", handlerScrapeFeeds)
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))
	return cmds
}

func getNewCommands() commands {
	cmdMap := make(map[string]func(*state, command) error)
	return commands{cmd_names: cmdMap}
}

func (c *commands) run(s *state, cmd command) error {
	run, exists := c.cmd_names[cmd.name]
	if !exists {
		return errors.New("command does not exist")
	}
	err := run(s, cmd)
	if err != nil {
		return errors.New("command failed: " + err.Error())
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd_names[name] = f
}
