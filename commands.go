package main

import (
	"context"
	"fmt"

	"github.com/ChernakovEgor/gator/internal/config"
	"github.com/ChernakovEgor/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
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

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.User)
		if err != nil {
			return fmt.Errorf("user not logged in: %v", err)
		}
		err = handler(s, cmd, user)
		return err
	}
}
