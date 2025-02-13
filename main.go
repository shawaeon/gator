package main

import (
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
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", handlerFollowFeed)

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