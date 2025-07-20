package main

import (
	"errors"
	"fmt"
)

// handlerLogin func is for login cmd, it sets the given user to the config
func handlerLogin(st *state, cmd command) error {
	if cmd.Arg == nil {
		return errors.New("handlerLogin expects a single argument, the username")
	}
	username := cmd.Arg[0]
	if err := st.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("user with the name of %v has been set\n", st.cfg.CurrentUsername)
	return nil
}