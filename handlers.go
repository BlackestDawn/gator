package main

func registerHandlers(c *commands) {
	c.register("login", handlerLogin)

}
