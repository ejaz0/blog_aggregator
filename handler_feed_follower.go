package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ejaz0/blog_aggreator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollower(s *state, cmd command, u database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("takes in url as argument")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("get feed: %w", err)
	}

	followFeed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    feed.Feedid,
	})
	if err != nil {
		return fmt.Errorf("getting follow feed: %v", err)
	}
	fmt.Printf("User %s is now following Feed %s\n", u.Name, followFeed.FeedName)
	return nil

}
