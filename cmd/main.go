package main

import (
	"log"

	"example.com/playground-wire-app/cmd/di"
)

func main() {
	app, cleanup := di.InitializeApp()

	err := app.Start()
	if err != nil {
		cleanup()
		log.Fatalf("app got error: %v", err)
	}
}
