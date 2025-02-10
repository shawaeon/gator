package main

import (
	"fmt"
	"log"

	"github.com/shawaeon/gator/internal/config"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	err = config.SetUser(&cfg, "shawaeon")
	if err != nil {
		log.Fatalf("error reafing config: %v",err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Println(cfg)
}