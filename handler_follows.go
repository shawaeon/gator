package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shawaeon/gator/internal/database"
)

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}

	ctx := context.Background()
	url := cmd.Args[0]
	
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
	fmt.Printf("* Feedname: %s\n", insertedFeedFollow.FeedName)
	fmt.Printf("* Username: %s\n", insertedFeedFollow.UserName)
	fmt.Println()
	fmt.Println("===============================================================")
	return nil
}

func handlerGetFeedFollows(s *state, cmd command, user database.User) error {
	fetchedFeedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not fetch follows: %w", err)
	}

	if len(fetchedFeedFollows) == 0 {
		println("You are not following any feeds. Use follow <RSS feed url>.")
		return nil
	}

	fmt.Println("Following feeds:")
	for _, follow := range fetchedFeedFollows {
		fmt.Printf("* Feedname: %s\n", follow.FeedName)
	}
	fmt.Println()
	fmt.Printf("* Username: %s\n", user.Name)
	fmt.Println()
	fmt.Println("===============================================================")
	
	return nil
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}
	ctx := context.Background()
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("could not find feed: %w", err)
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not find follow: %w", err)
	}
	
	fmt.Println("Unfollowed:")
	fmt.Printf("* Feed: %s\n", feed.Name)
	fmt.Printf("* User: %s\n", user.Name)
	fmt.Println()
	fmt.Println("===============================================================")


	return nil
}