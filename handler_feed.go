package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/h0dy/blog-aggregator/internal/database"
)

func handlerAddFeed(st *state, cmd command, user database.User) error {
	if len(cmd.Arg) < 2 {
		return fmt.Errorf("\nusage: %s <feed name> <feed url>", cmd.Name)
	}

	feedName := cmd.Arg[0]
	feedURL := cmd.Arg[1]
	feed, err := st.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: 	   uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: feedName,
		Url: feedURL,
		UserID:user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %v", err)
	}

	feedFollow, err := st.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: 	   uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't follow current feed: %v", err)
	}

	fmt.Println("Feed created successfully")
	logFeed(feed)
	fmt.Println("Feed followed successfully")
	printFeedFollow(feedFollow, feed.Url)
	return nil
}

func handlerListFeeds(st *state, cmd command) error {
	feeds, err := st.db.GetAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("\ncouldn't get feeds: %v", err)
	}
	if len(feeds) < 1 {
		fmt.Println("No feeds found ☹️\n use addfeed <name> <url> command to add feeds.")
		return nil
	}
	
	logFeeds(feeds)
	return nil
}

func logFeeds(feeds []database.GetAllFeedsRow) {
	for _, feed := range feeds {
		fmt.Println("===============================")
		fmt.Printf("Feed name: %v\n", feed.FeedName)
		fmt.Printf("Feed URL: %v\n", feed.FeedUrl)
		fmt.Printf("Created by: %v\n", feed.CreatedBy)
	}
}

func logFeed(feed database.Feed) {
	fmt.Printf("Feed details:\nFeed ID: %v\nFeed Name: %v\nFeed URL: %v\n", 
	feed.ID,
	feed.Name,
	feed.Url,
	)
}
