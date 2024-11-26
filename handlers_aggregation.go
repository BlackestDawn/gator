package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/BlackestDawn/gator/internal/database"
)

func handlerAggregation(s *state, cmd command) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("usage: gator agg <time interval>")
	}

	interval, err := time.ParseDuration(cmd.params[0])
	if err != nil {
		return fmt.Errorf("error parsing time interval: %w", err)
	}

	fmt.Println("Collecting feeds every", interval)
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feed: %w", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background(), s.conf.CurrentUserName)
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	cTime := sql.NullTime{Time: time.Now(), Valid: true}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{LastFetchedAt: cTime, UpdatedAt: time.Now(), ID: nextFeed.ID})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	for _, item := range rssFeed.Channel.Item {
		err := createPost(s, item, nextFeed.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func printFeed(feed *RSSFeed) {
	fmt.Println(feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println("  -", item.Title)
	}
	fmt.Println("")
}
