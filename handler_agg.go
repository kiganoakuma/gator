package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kiganoakuma/gator/internal/database"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) > 1 {
		log.Fatalf("usage: %s <time_between_reqs>", cmd.Name)
	}
	cmdTime := cmd.Args[0]
	if len(cmd.Args) == 0 {
		cmdTime = "5m"
	}

	timeBetweenReqs, err := time.ParseDuration(cmdTime)
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %v", err)
	}

	currentTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		LastFetchedAt: currentTime,
		UpdatedAt:     time.Now(),
	})
	if err != nil {
		return err
	}

	nextRssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, feed := range nextRssFeed.Channel.Item {
		description := sql.NullString{
			String: feed.Description,
			Valid:  feed.Description != "",
		}

		pubDate := parsePublishedDate(feed.PubDate)

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       feed.Title,
			Description: description,
			PublishedAt: pubDate,
			Url:         feed.Link,
			FeedID:      nextFeed.ID,
		})

		if err != nil {
			// Check if it's a PostgreSQL error
			if pqErr, ok := err.(*pq.Error); ok {
				// Check if it's a duplicate key error (error code 23505)
				if pqErr.Code == "23505" && strings.Contains(pqErr.Constraint, "posts_url_key") {
					// Silently ignore this specific error (post with URL already exists)
					continue
				}
			}
			// Log any other errors
			log.Printf("Error creating post: %v", err)
		}
	}
	return nil
}

func parsePublishedDate(pubDate string) time.Time {
	// Try the most common RSS date format first (RFC1123Z)
	parsedTime, err := time.Parse(time.RFC1123Z, pubDate)
	if err == nil {
		return parsedTime
	}

	// If the common format failed, try the other formats
	formats := []string{
		time.RFC1123,                // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,                // "02 Jan 06 15:04 -0700"
		time.RFC822,                 // "02 Jan 06 15:04 MST"
		"2006-01-02T15:04:05Z07:00", // ISO 8601 / RFC3339
	}

	for _, format := range formats {
		parsedTime, err = time.Parse(format, pubDate)
		if err == nil {
			return parsedTime
		}
	}

	// If all parsing attempts fail, return current time
	return time.Now()
}
