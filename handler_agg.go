package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/h0dy/blog-aggregator/internal/database"
)

func handlerFeed(st *state, cmd command)error {
	if len(cmd.Arg) < 1 {
		return fmt.Errorf("\nusage %v <time between requests>", cmd.Name)
	}
	timeBetweenRqs, err := time.ParseDuration(cmd.Arg[0])
	if err != nil {
		return err
	}

	fmt.Printf("collecting feeds every %v...\n", timeBetweenRqs)

	ticker := time.NewTicker(timeBetweenRqs)
	defer ticker.Stop()

	for ;; <-ticker.C {
		scrapeFeeds(st)
	}
}


func scrapeFeeds(st *state) {
	nextFeed, err := st.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("\nCouldn't get feeds to fetch: %v", err)
		return
	}

	log.Printf("Found a feed to fetch")
	scrapeFeed(st.db, nextFeed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	if err := db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		fmt.Printf("\ncouldn't mark the feed as fetched: %v", err)
		return
	}
	
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("\ncouldn't collect feed %v: %v", feed.Name, err)
		return
	}
	
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Feed channel title: %v\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}