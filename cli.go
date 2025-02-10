package main

import (
	"fmt"

	"github.com/shawaeon/gator/internal/config"
)


type command struct {
	name 		string
	arguments	[]string
}

type commands struct {
	commandNames	map[string]func(*state, command) error
}

// Add a function to commands
func (c *commands) register(name string, f func(*state, command) error) {
	c.commandNames[name] = f
}

// Run a function from commans
func (c *commands) run(s *state, cmd command) error {
	commandFunc, ok := c.commandNames[cmd.name]
	if !ok {
		return fmt.Errorf("error: command not found")
	}
	err := commandFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

// Login as cmd.arguments[0]
func handlerLogin (s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("error: no username")
	}
	username := cmd.arguments[0]

	err := config.SetUser(s.cfg, username)
	if err != nil {
		return err
	}

	fmt.Printf("Logged in as %s\n", username)
	return nil
}