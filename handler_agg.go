package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/h0dy/blog-aggregator/internal/database"
)

func handlerFeed(st *state, cmd command)error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerAddFeed(st *state, cmd command) error {
	if len(cmd.Arg) < 2 {
		return fmt.Errorf("\nusage: %s <feed name> <feed url>", cmd.Name)
	}
	user, err := st.db.GetUser(context.Background(), st.cfg.CurrentUsername)
	if err != nil {
		return err
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
	fmt.Println("Feed created successfully")
	logFeed(feed)
	return nil
}

func logFeed(feed database.Feed) {
	fmt.Printf("FEED\nFeed ID: %v\nFeed Name: %v\nFeed URL: %v\nCreated at: %v\nUpdated at: %v\nUser ID: %v\n", 
	feed.ID,
	feed.Name,
	feed.Url,
	feed.CreatedAt,
	feed.UpdatedAt,
	feed.UserID,
	)
}