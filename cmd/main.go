package main

import (
	"log"

	"example.com/playground-wire-app/cmd/di"
)

func main() {
	app := di.InitializeApp()

	err := app.Start()
	if err != nil {
		log.Fatalf("app got error: %v", err)
	}
}
