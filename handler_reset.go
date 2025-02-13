package main

import (
	"context"
	"fmt"
)

func handlerReset (s *state, cmd command) error {
		
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database: %w", err)
	}
	fmt.Println("All users removed")
	return nil
}