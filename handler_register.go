package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shawaeon/gator/internal/config"
	"github.com/shawaeon/gator/internal/database"
)

// Add a new user to the database
func handlerRegister (s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	uuid := uuid.New()	
	timeNow := time.Now()
	username := cmd.Args[0]

	insertedUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name: username,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	
	err = config.SetUser(s.cfg, username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User added: \n%v\n", insertedUser.Name)

	return nil
}