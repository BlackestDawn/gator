package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlackestDawn/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("usage: gator login <name>")
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

func handlerRegister(s *state, cmd command) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("usage: gator register <name>")
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

func handlerListUsers(s *state, cmd command) error {
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
