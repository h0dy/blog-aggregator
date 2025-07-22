package main

import (
	"context"
	"fmt"
)

func handlerReset(st *state, cmd command)error {
	if err := st.db.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("\nerror in db.DeleteAllUsers: %w", err)
	}
	fmt.Println("Users table reset successfully")
	return nil
}
