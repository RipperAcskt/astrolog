package main

import (
	"log"

	"astrolog/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
