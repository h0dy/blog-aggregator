package main

import (
	"context"
	"fmt"
)

func handlerFeed(st *state, cmd command)error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
