package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/BlackestDawn/gator/internal/database"
	"github.com/google/uuid"
)

func createPost(s *state, feedItem RSSItem, feedID uuid.UUID) error {
	cTime := time.Now()
	parsedTime := sql.NullTime{}
	if t, err := time.Parse(time.RFC1123Z, feedItem.PubDate); err == nil {
		parsedTime = sql.NullTime{
			Time:  t,
			Valid: true,
		}
	}

	data := database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: cTime,
		UpdatedAt: cTime,
		Title:     feedItem.Title,
		Url:       feedItem.Link,
		Description: sql.NullString{
			String: feedItem.Description,
			Valid:  true,
		},
		PublishedAt: parsedTime,
		FeedID:      feedID,
	}
	err := s.db.CreatePost(context.Background(), data)
	if err != nil && !strings.Contains(err.Error(), "duplicate key value") {
		return fmt.Errorf("error saving post: %w", err)
	}

	return nil
}

func handlerGetPosts(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.params) >= 1 {
		if tmp, err := strconv.Atoi(cmd.params[0]); err == nil {
			limit = tmp
		} else {
			return fmt.Errorf("error setting limit: %w", err)
		}
	}

	data := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), data)
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}

	for _, item := range posts {
		fmt.Println("ID:", item.ID)
		fmt.Println("Created at:", item.CreatedAt)
		fmt.Println("Updated at:", item.UpdatedAt)
		fmt.Println("Title:", item.Title)
		fmt.Println("URL:", item.Url)
		fmt.Println("Description:", item.Description.String)
		fmt.Println("Published at:", item.PublishedAt.Time)
		fmt.Println("Feed ID:", item.FeedID)
		fmt.Println("")
	}

	return nil
}
