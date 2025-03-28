package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kiganoakuma/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feedFollowRec, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create feed follow record: %w", err)
	}

	fmt.Println(feedFollowRec.UserName, feedFollowRec.FeedName)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s <user_name>", cmd.Name)
	}
	userFollowing, err := s.db.GetFeedFollowesForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't retrieve users feed: %w", err)
	}
	for _, rec := range userFollowing {
		fmt.Println(rec.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("useage: %s <feed_url>", cmd.Name)
	}

	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
