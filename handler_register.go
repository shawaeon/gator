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
	username := cmd.Args[0]

	insertedUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	
	err = config.SetUser(s.cfg, username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User added:")
	printUser(insertedUser)

	return nil
}

func printUser(user database.User) {
	fmt.Printf("* ID:			%v\n", user.ID)
	fmt.Printf("* Name:			%s\n", user.Name)
	fmt.Println()
	fmt.Println("===============================================================")
}