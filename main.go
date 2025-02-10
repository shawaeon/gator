package main

import (
	"log"

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
	loginCommand := command {
		name: "login",
		arguments: []string{"test"},
	}

	testCommands := commands {
		commandNames: map[string]func(*state, command) error {},
	}
	testCommands.register(loginCommand.name, handlerLogin)

	err = testCommands.run(&appState, loginCommand)
	if err != nil {
		log.Fatalf("error logging in: %v", err)
	}
}