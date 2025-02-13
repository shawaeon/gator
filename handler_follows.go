package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shawaeon/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}

	ctx := context.Background()
	url := cmd.Args[0]
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	
	fetchedFeed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)		
	} 
	
	insertedFeedFollow, err :=  s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: fetchedFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create follow: %w", err)
	}

	fmt.Println("Following:")
	fmt.Printf("* Feed: %s\n", insertedFeedFollow.FeedName)
	fmt.Printf("* User: %s\n", insertedFeedFollow.UserName)
	fmt.Println()
	fmt.Println("===============================================================")
	return nil
}