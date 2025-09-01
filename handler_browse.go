package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ejaz0/blog_aggreator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	// parse optional limit from cmd.Args, default 2
	limit := int32(2)
	if len(cmd.Args) >= 1 {
		if n, err := strconv.Atoi(cmd.Args[0]); err == nil && n > 0 {
			limit = int32(n)
		}
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, p := range posts {
		fmt.Printf("%s\n%s\n%s\n\n", p.Title, p.Url, p.PublishedAt.Format(time.RFC3339))
	}
	return nil
}
