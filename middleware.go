package main

import (
	"context"
	"fmt"

	"github.com/ejaz0/blog_aggreator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		u, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("load current user: %w", err)
		}
		err = handler(s, cmd, u)
		if err != nil {
			return err
		}
		return nil
	}
}
