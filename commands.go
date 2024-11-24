package main

import "errors"

type command struct {
	name   string
	params []string
}

type commands struct {
	registered map[string]func(*state, command) error
}

func NewCommands() *commands {
	cmds := new(commands)
	cmds.registered = map[string]func(*state, command) error{}
	registerHandlers(cmds)
	return cmds
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registered[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	reg, exists := c.registered[cmd.name]
	if !exists {
		return errors.New("command not found")
	}
	return reg(s, cmd)
}
