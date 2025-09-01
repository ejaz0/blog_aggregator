package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ejaz0/blog_aggreator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, u database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("needs two arguments")
	}
	name, url := cmd.Args[0], cmd.Args[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    u.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
	}

	f, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("create feed %w", err)
	}
	followPrams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    u.ID,
		FeedID:    f.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followPrams)
	if err != nil {
		return fmt.Errorf("vreating feed follow: %v", err)
	}
	fmt.Println("ID:", f.ID)
	fmt.Println("UserID:", f.UserID)
	fmt.Println("Name:", f.Name)
	fmt.Println("URL:", f.Url)
	fmt.Println("CreatedAt:", f.CreatedAt)
	fmt.Println("UpdatedAt:", f.UpdatedAt)
	return nil

}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("get feeds: %w", err)
	}
	for _, f := range feeds {
		fmt.Printf("Feed Name: %v\n", f.FeedName)
		fmt.Printf("Feed Name: %v\n", f.Url)
		fmt.Printf("Feed Name: %v\n", f.UserName)
	}
	return nil
}
