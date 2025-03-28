package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kiganoakuma/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <limit> default: %d", cmd.Name, limit)
	}
	if len(cmd.Args) == 1 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("limit must be a valid integer: %v", err)
		}
		limit = parsedLimit
	}

	userPosts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldnt retrieve posts from user: %w", err)
	}

	for _, post := range userPosts {
		fmt.Printf("title: %s\n", post.Title)
		fmt.Printf("user_id: %d\n", post.UserID)
	}

	return nil
}
