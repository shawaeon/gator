package main

import (
	"log"

	"github.com/shawaeon/gator/internal/config"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	s := state {
		cfg: &cfg,
	}
	loginCommand := command {
		name: "login",
		arguments: []string{"test"},
	}
	handlerLogin(&s, loginCommand)
}