package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shawaeon/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	
	return nil
}