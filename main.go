package main

import (
	"fmt"
	"log"
	"os"

	"github.com/h0dy/blog-aggregator/internal/config"
)

// state struct holds a pointer to a config
type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.ReadConfigFile()
	if err != nil {
		fmt.Printf("error reading config file: %v\n", err)
	}

	userState := &state{
		cfg: &cfg,
	}

	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	
	commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments ")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("error: username is required for login command")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArg := os.Args[2:]

	if err = commands.run(userState, command{
		Name:cmdName,
		Arg: cmdArg,
	}); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("config file: %v\n", cfg)
}