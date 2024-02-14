package main

import (
	"log"
	"os"

	"example.com/playground-wire-app/cmd/di"
	"example.com/playground-wire-app/internal/app"
)

func initApp(env string) (*app.App, func(), error) {
	switch env {
	case "DEBUG", "debug":
		return di.InitializeAppDebug()
	}

	return di.InitializeApp()
}

func main() {
	envENV, _ := os.LookupEnv("ENV")

	app, cleanup, err := initApp(envENV)
	if err != nil {
		cleanup()
		log.Fatalf("failed to init app: %v", err)
	}

	err = app.Start()
	if err != nil {
		cleanup()
		log.Fatalf("app got error: %v", err)
	}
}
