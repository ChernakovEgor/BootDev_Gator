package main

import (
	"context"
	"fmt"
)

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
