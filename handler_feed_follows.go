package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/h0dy/blog-aggregator/internal/database"
)

func handlerFollowFeed(st *state, cmd command) error {
	if len(cmd.Arg) < 1 {
		return fmt.Errorf("\nusage: %s <feed url>", cmd.Name)
	}

	feedURL := cmd.Arg[0]
	feed, err := st.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("\ncouldn't get feed by url: %w", err)
	}

	user, err := st.db.GetUser(context.Background(), st.cfg.CurrentUsername)
	if err != nil {
		return err
	}

	feedFollowed, err := st.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: 	   uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("\ncouldn't follow the feed: %w", err)
	}

	printFeedFollow(feedFollowed, feed.Url)
	return nil
}

func handlerFollowingFeeds(st *state, cmd command)error {
	followFeeds, err := st.db.GetFeedFollowsForUser(context.Background(), st.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("\ncouldn't get user's following feeds: %v", err)
	}

	fmt.Printf("%v follows:\n", st.cfg.CurrentUsername)
	for _, f := range followFeeds {
		fmt.Println("===========================")
		fmt.Printf("Feed name: %v\n", f.FeedName)
		fmt.Printf("Feed URL: %v\n", f.FeedUrl)
	}
	
	return nil
}

func printFeedFollow(feedFollowed database.CreateFeedFollowRow, feedUrl string) {
	fmt.Printf("USER:\n")
	fmt.Printf("User ID: %v\n", feedFollowed.UserID)
	fmt.Printf("User name: %v\n", feedFollowed.UserName)
	fmt.Printf("Follow \"%v\"\nFeed URL: %v\n", feedFollowed.FeedName, feedUrl)
}