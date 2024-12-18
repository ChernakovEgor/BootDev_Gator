package main

import (
	"fmt"

	"github.com/ChernakovEgor/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdName := cmd.name
	handler, ok := c.cmds[cmdName]
	if !ok {
		return fmt.Errorf("command not found")
	}

	err := handler(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("incorrect number of arguments")
	}

	userName := cmd.args[0]
	err := s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", userName)

	return nil
}
