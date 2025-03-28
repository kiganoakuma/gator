package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kiganoakuma/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couln't create feed: %w", err)
	}
	if _, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}); err != nil {
		return fmt.Errorf("couldnt insert feed record: %w", err)
	}

	fmt.Println("Feed created successfully")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("fetching feeds table failed: %w", err)
	}

	for _, feed := range feeds {
		userName, err := s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't fetch username: %w", err)
		}
		fmt.Println(feed.Name)
		fmt.Println(userName)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:               %s\n", feed.ID)
	fmt.Printf("* Created:          %s\n", feed.CreatedAt)
	fmt.Printf("* Updated:          %s\n", feed.UpdatedAt)
	fmt.Printf("* Name:             %s\n", feed.Name)
	fmt.Printf("* URL:              %s\n", feed.Url)
	fmt.Printf("* UserId:           %s\n", feed.UserID)
}
