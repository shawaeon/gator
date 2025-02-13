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
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed: %w", err)
	}

	// Follow created feed
	handlerFollowFeed(s, command {
		Name: "follow",
		Args: []string{insertedFeed.Url},
	})

	fmt.Println("Feed created successfully:")
	printFeed(insertedFeed)
	fmt.Println()
	fmt.Println("===============================================================")

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	ctx := context.Background()
	fetchedFeeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds: %w", err)
	}

	if len(fetchedFeeds) == 0 {
		println("No feeds found. Use addfeed <RSS feed name> <RSS feed url>.")
		return nil
	}

	fmt.Println("Your feeds:")
	for _, feed := range fetchedFeeds {
		userName, err := getUserName(s, feed)
		if err != nil {
			return fmt.Errorf("could not fetch user associated with feed: %w", err)
		}
		
		printFeed(feed)
		fmt.Printf("* Username:		%s\n", userName)
		fmt.Println()
	}
	fmt.Println("===============================================================")

	return nil
}

func getUserName(s *state, feed database.Feed) (string, error) {
	ctx := context.Background()
	fetchedUser, err := s.db.GetUserByID(ctx, feed.UserID)
	if err != nil {
		return "", err
	}
	return fetchedUser.Name, nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:			%s\n", feed.ID)
	fmt.Printf("* CreatedAt:		%s\n", feed.CreatedAt)
	fmt.Printf("* UpdatedAt:		%s\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
	fmt.Printf("* UserID:		%s\n", feed.UserID)
}