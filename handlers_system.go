package main

import (
	"context"
	"fmt"
)

func registerHandlers(c *commands) {
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerListUsers)
	c.register("agg", handlerAggregation)
	c.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	c.register("feeds", handlerListFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", handlerListFollowing)
	c.register("unfollow", middlewareLoggedIn(handlerDeleteFollow))
	c.register("browse", middlewareLoggedIn(handlerGetPosts))
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not clear users table: %w", err)
	}

	return nil
}
