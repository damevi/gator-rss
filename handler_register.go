package main

import (
	"context"
	"fmt"
	"time"

	"github.com/damevi/gator-rss/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})

	if err != nil {
		return fmt.Errorf("couldn't create new user: %w", err)
	}

	fmt.Println("User created succcessfully!")
	fmt.Printf("Created user data: %v\n", user)

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	return nil
}
