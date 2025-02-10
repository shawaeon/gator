package main

import (
	"fmt"
)

type command struct {
	Name 		string
	Args	[]string
}

type commands struct {
	registeredCommands	map[string]func(*state, command) error
}

// Add a function to commands
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// Run a function from commans
func (c *commands) run(s *state, cmd command) error {
	commandFunc, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("error: command not found")
	}
	err := commandFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
