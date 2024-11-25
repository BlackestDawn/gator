package main

import (
	"context"
)

func registerHandlers(c *commands) {
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerListUsers)
	c.register("agg", handlerAggregation)
	c.register("addfeed", handlerAddfeed)
	c.register("feeds", handlerListFeeds)
}

func handlerReset(s *state, cmd command) error {
	return s.db.ResetUsers(context.Background())
}
