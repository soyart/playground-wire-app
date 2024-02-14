package app

import (
	"log"

	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/repo"
)

type App struct {
	conf config.Config
	repo repo.Repo
}

func ProvideApp(conf config.Config, repo repo.Repo) App {
	return App{
		conf: conf,
		repo: repo,
	}
}

func (a *App) Start() error {
	log.Println("app start")

	log.Println("app shutting down")
	return nil
}