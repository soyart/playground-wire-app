package main

import (
	"log"
	"os"

	"example.com/playground-wire-app/cmd/di"
	"example.com/playground-wire-app/internal/app"
)

func initApp(env string) (*app.App, func()) {
	switch env {
	case "DEBUG", "debug":
		return di.InitializeAppDebug()
	}

	return di.InitializeApp()
}

func main() {
	envENV, _ := os.LookupEnv("ENV")
	app, cleanup := initApp(envENV)

	err := app.Start()
	if err != nil {
		cleanup()
		log.Fatalf("app got error: %v", err)
	}
}
