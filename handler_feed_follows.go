package main

import (
	"context"
	"fmt"

	"github.com/ejaz0/blog_aggreator/internal/database"
)

func handlerFeedFollows(s *state, cmd command, u database.User) error {

	// Print user's name here, only once
	fmt.Printf("Feeds followed by %s:\n", u.Name) // Example of printing it once
	fmt.Println("--------------------")

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), u.ID)
	if err != nil {
		return fmt.Errorf("getting all feed follows: %v", err)
	}

	for _, f := range follows {
		fmt.Printf("  - %v\n", f.FeedName) // Added a newline and some formatting
	}

	// Don't forget to return nil if everything is successful!
	return nil
}
