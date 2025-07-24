package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/h0dy/blog-aggregator/internal/database"
)

func handlerPosts(st *state, cmd command, user database.User) error {
	var postLimit int32 = 2
	if len(cmd.Arg) > 1 {
		limit, err := strconv.Atoi(cmd.Arg[0])
		if err != nil {
			return err
		}
		postLimit = int32(limit)
	}
	posts, err := st.db.GetPosts(context.Background(), database.GetPostsParams{
		UserID: user.ID,
		Limit: postLimit,
	})
	if len(posts) < 1 {
		fmt.Println("There are no posts for user: %v", user.Name)
	}

	if err != nil {
		return fmt.Errorf("\ncouldn't get posts: %v", err)
	}
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}