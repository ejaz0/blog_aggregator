package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ejaz0/blog_aggreator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("get next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("marking feed fetched: %w", err)
	}

	rss, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("parsing feed %w", err)
	}
	parsePub := func(s string) (time.Time, error) {
		for _, layout := range []string{time.RFC1123Z, time.RFC1123, time.RFC822Z, time.RFC822, time.RFC3339} {
			if t, e := time.Parse(layout, s); e == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("unparsed pubDate")
	}

	now := time.Now()

	for _, it := range rss.Channel.Item {
		pubAt, perr := parsePub(it.PubDate)
		if perr != nil {
			continue
		}
		desc := sql.NullString{String: it.Description, Valid: it.Description != ""}
		p := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       it.Title,
			Url:         it.Link,
			Description: desc,
			PublishedAt: pubAt,
			FeedID:      feed.ID,
		}
		if err = s.db.CreatePost(ctx, p); err != nil {
			fmt.Printf("create post error for %q: %v\n", it.Link, err)
		}
	}

	return nil
}
