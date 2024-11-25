package main

import (
	"context"
	"fmt"
)

func handlerAggregation(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item)
	}

	return nil
}
