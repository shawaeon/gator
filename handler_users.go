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
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	
	username := cmd.Args[0]

	insertedUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
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

// Login as cmd.arguments[0]
func handlerLogin (s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}	
	
	username := cmd.Args[0]

	fetchedUser, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = config.SetUser(s.cfg, username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Logged in as %s\n", fetchedUser.Name)

	return nil
}

func printUser(user database.User) {
	fmt.Printf("* ID:			%v\n", user.ID)
	fmt.Printf("* Name:			%s\n", user.Name)
	fmt.Println()
	fmt.Println("===============================================================")
}

func handlerListUsers(s *state, cmd command) error {
	fetchedUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch users, %w", err)
	}

	if len(fetchedUsers) == 0 {
		fmt.Println("Couldn't find any users. Use register <name>")
		return nil
	}
	fmt.Println("Users:")
	for _, user := range(fetchedUsers) {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	
	return nil
}