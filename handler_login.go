package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("no username given")
	}

	name := cmd.params[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("could not find user: %w", err)
	}

	err = s.conf.SetUser(name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}
	fmt.Println("Successfully set user to:", s.conf.CurrentUserName)
	return nil
}
