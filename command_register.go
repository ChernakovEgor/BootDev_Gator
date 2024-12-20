package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ChernakovEgor/gator/internal/database"
	"github.com/google/uuid"
)

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
