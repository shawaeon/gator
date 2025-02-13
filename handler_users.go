package main

import (
	"context"
	"fmt"
)

func handlerListUsers(s *state, cmd command) error {
	fetchedUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch users, %w", err)
	}

	if len(fetchedUsers) == 0 {
		fmt.Println("Couldn't find any users. Use register <name>")
		return nil
	}
	fmt.Println("Users:")
	for _, user := range(fetchedUsers) {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	
	return nil
}