package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ChernakovEgor/gator/internal/config"
	"github.com/ChernakovEgor/gator/internal/database"
	"github.com/google/uuid"
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("incorrect number of arguments")
	}

	userName := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", userName)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("not enough arguments")
	}

	userName := cmd.args[0]
	userParams := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: userName}
	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", userName)

	return nil
}

func handlerReset(s *state, _ command) error {
	err := s.db.Reset(context.Background())
	return err
}

func handlerUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %v", err)
	}

	for _, user := range users {
		if s.cfg.User == user.Name {
			fmt.Printf(" * %s (current)\n", user.Name)
		} else {
			fmt.Printf(" * %s\n", user.Name)
		}
	}
	return nil
}
