package main

import (
	"context"
	"fmt"
)

func handlerReset (s *state, cmd command) error {
	ctx := context.Background()
	
	err := s.db.ResetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error resetting database: %w", err)
	}
	fmt.Println("All users removed")
	return nil
}