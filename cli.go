package main

import (
	"fmt"

	"github.com/shawaeon/gator/internal/config"
)

type state struct {
	cfg		*config.Config
}

type command struct {
	name 		string
	arguments	[]string
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