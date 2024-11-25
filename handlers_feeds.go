package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlackestDawn/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.params) < 2 {
		return fmt.Errorf("usage: gator addfeed <name> <url>")
	}

	cUser, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get user info: %w", err)
	}

	name := cmd.params[0]
	url := cmd.params[1]
	cTime := time.Now()
	newFeed := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: cTime,
		UpdatedAt: cTime,
		Name:      name,
		Url:       url,
		UserID:    cUser.ID,
	}
	_, err = s.db.AddFeed(context.Background(), newFeed)
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}

	fmt.Println("Successfully added feed:")
	fmt.Println(newFeed)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	feedList, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch feeds: %w", err)
	}

	fmt.Println("List of current feeds:")
	fmt.Println("    Name:    Url:     Owner:")
	for _, item := range feedList {
		fmt.Printf("  * %v: %v - %v\n", item.Name, item.Url, item.Username)
	}
	return nil
}
