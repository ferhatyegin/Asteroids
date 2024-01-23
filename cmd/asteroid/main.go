package main

import (
	"log"

	asteroid "github.com/ferhatyegin/goAsteroids"
)

func main() {
	game := asteroid.NewGame()

	if err := game.Run(); err != nil {
		log.Fatalf("Game Error : %v", err)
	}
}
