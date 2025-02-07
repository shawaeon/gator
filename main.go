package main

import (
	"fmt"

	"github.com/shawaeon/gator/internal/config"
)

func main(){
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Testing config:")
	fmt.Println(config.DbURL)
}