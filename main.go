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

	if len(os.Args) < 2 {
		fmt.Println("error: no command arguments")
		os.Exit(1)
	}

	commandInput := os.Args[1]
	commandArgs := os.Args[2:]

	testCommand := command {
		name: commandInput,
		arguments: commandArgs,
	}

	testCommands := commands {
		commandNames: map[string]func(*state, command) error {},
	}
	testCommands.register(testCommand.name, handlerLogin)

	err = testCommands.run(&appState, testCommand)
	if err != nil {
		log.Fatal(err)
	}
}