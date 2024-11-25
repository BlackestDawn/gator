package main

import (
	"context"
	"fmt"
	"time"

	"github.com/BlackestDawn/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("usage: gator follow <url>")
	}

	url := cmd.params[0]
	feedData, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feed data: %w", err)
	}

	cTime := time.Now()
	data := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: cTime,
		UpdatedAt: cTime,
		UserID:    user.ID,
		FeedID:    feedData.ID,
	}
	retData, err := s.db.CreateFeedFollow(context.Background(), data)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Println("Successfully created follow request:")
	fmt.Println(retData)
	return nil
}

func handlerListFollowing(s *state, cmd command) error {
	feedData, err := s.db.GetFeedsFollowForUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching feed data: %w", err)
	}

	fmt.Println("You are following these feeds:")
	for _, item := range feedData {
		fmt.Println("  -", item.FeedName)
	}

	return nil
}

func handlerDeleteFollow(s *state, cmd command, user database.User) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("usage: gator unfollow <url>")
	}

	url := cmd.params[0]
	feedData, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feed data: %w", err)
	}

	data := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feedData.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(), data)
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	return nil
}
