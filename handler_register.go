package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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
		return err
	}
	
	s.cfg.CurrentUserName = username
	fmt.Printf("User added: %v\n", username)
	log.Printf("User added: %v\n", insertedUser)

	return nil
}