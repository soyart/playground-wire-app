package app

import (
	"log"

	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/repo"
)

type App struct {
	Configuration config.Config
	Repository    repo.Repo
}

func (a *App) Start() error {
	log.Println("app start")

	log.Println("app shutting down")
	return nil
}
