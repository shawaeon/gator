package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shawaeon/gator/internal/config"
)


type state struct {
	cfg		*config.Config
}

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := state {
		cfg: &cfg,
	}

	cmds := commands {
		registeredCommands: map[string]func(*state, command) error {},
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	testCommand := command {
		Name: commandName,
		Args: commandArgs,
	}	

	err = cmds.run(&appState, testCommand)
	if err != nil {
		log.Fatal(err)
	}
}