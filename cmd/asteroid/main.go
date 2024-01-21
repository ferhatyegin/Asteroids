package main

import (
	"fmt"
	"log"

	asteroid "github.com/ferhatyegin/goAsteroids"
)

func main() {
	fmt.Println("Hello World")

	game := asteroid.NewGame()

	if err := game.Run(); err != nil {
		log.Fatalf("Game Error : %v", err)
	}
}