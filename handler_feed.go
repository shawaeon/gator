package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shawaeon/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <RSS feed name> <RSS feed url>", cmd.Name)
	}
	
	ctx := context.Background()	
	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	insertedFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(insertedFeed)
	fmt.Println()
	fmt.Println("===============================================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:			%s\n", feed.ID)
	fmt.Printf("* CreatedAt:		%s\n", feed.CreatedAt)
	fmt.Printf("* UpdatedAt:		%s\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
	fmt.Printf("* UserID:		%s\n", feed.UserID)
}