package main

import "errors"

// command struct hold the name of the command and its arguments
type command struct {
	Name string
	Arg []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

// run func runs a given command with the provided state if it exists
func (c *commands) run(st *state, cmd command) error {
	// get the command and check if it exists
	cmdFunc, ok := c.cmds[cmd.Name] 
	if !ok  {
		return errors.New("the given command isn't registered")
	}
	if err := cmdFunc(st, cmd); err != nil {
		return err
	} // run the command
	return nil
}

// register func registers a new handler function for a command name
func (c *commands) register(name string, fn func(*state, command) error) {
	c.cmds[name] = fn
}
