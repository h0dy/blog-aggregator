package main

import (
	"fmt"

	"github.com/h0dy/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.ReadConfigFile()
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
	}
	
	if err := cfg.SetUser("hody"); err != nil {
		fmt.Printf("error setting the user: %v\n", err)
	}

	fmt.Println(cfg)
}