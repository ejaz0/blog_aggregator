package main

import (
	"context"
	"fmt"

	"github.com/ejaz0/blog_aggreator/internal/database"
)

func handlerUnfollow(s *state, cmd command, u database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("takes in a url")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("getting feed: %w", err)
	}
	_, err = s.db.DeleteFeedFollowForUser(context.Background(), database.DeleteFeedFollowForUserParams{
		ID:   u.ID,        // $1 - user ID
		ID_2: feed.Feedid, // $2 - feed ID
	})
	if err != nil {
		return fmt.Errorf("deleting feed: %w", err)
	}
	fmt.Printf("Successfully unfollowed feed!\n")
	return nil
}
