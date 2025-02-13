package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/shawaeon/gator/internal/config"
	"github.com/shawaeon/gator/internal/database"
)


type state struct {
	db		*database.Queries
	cfg		*config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	appState := state {
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := commands {
		registeredCommands: map[string]func(*state, command) error {},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollowFeed))
	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollows))

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := command {
		Name: commandName,
		Args: commandArgs,
	}	

	err = cmds.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

// Check that user is logged in
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(s *state, cmd command) error {
	return func(s *state, cmd command) error{
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("not logged in")
		}
		return handler(s, cmd, user)
	}
}