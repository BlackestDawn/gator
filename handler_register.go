package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlackestDawn/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("no name given")
	}

	name := cmd.params[0]
	cTime := time.Now()
	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: cTime,
		UpdatedAt: cTime,
		Name:      name,
	}

	retUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.conf.SetUser(retUser.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Println("Successfully created user:")
	fmt.Println(retUser.Name)
	return nil
}
