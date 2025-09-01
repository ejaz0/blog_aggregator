package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ejaz0/blog_aggreator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err == sql.ErrNoRows {
		return fmt.Errorf("user %q does not exist", name)
	}
	if err != nil {
		return fmt.Errorf("lookup failed: %w", err)
	}
	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("user switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("not username passed in")
	}
	name := cmd.Args[0]
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		return fmt.Errorf("user already exists")
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("lookup failed %w", err)
	}
	id := uuid.New()
	now := time.Now()
	u, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("create failed: %w", err)
	}
	if err := s.cfg.SetUser(name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("created user %q\n", name)
	log.Printf("user: %+v\n", u)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil
}
