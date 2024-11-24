package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	names, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch users: %w", err)
	}

	for _, name := range names {
		if name != s.conf.CurrentUserName {
			fmt.Println("*", name)
		} else {
			fmt.Println("*", name, "(current)")
		}
	}

	return nil
}
