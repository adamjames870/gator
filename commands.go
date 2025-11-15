package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	cmd_names map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) {
	run, exists := c.cmd_names[cmd.name]
	if !exists {
		return
	}
	run(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd_names[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		// Empty or nil arguments, login expects a login name
		return errors.New("no login name provided")
	}
	s.config.Current_user_name = cmd.args[0]
	fmt.Printf("Set username to '%s'\n", s.config.Current_user_name)
	return nil
}
