package main

import (
	"fmt"

	"github.com/shawaeon/gator/internal/config"
)

// Login as cmd.arguments[0]
func handlerLogin (s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	err := config.SetUser(s.cfg, username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Logged in as %s\n", username)
	return nil
}