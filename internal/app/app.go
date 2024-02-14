package app

import (
	"errors"
	"log"
	"time"

	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/repo"
)

type App struct {
	Name          string
	Configuration config.Config
	Repository    repo.Repo
}

func (a *App) Start() error {
	if a.Name == "" {
		return errors.New("app has no name")
	}

	log.Printf("[%s] app start", a.Name)
	time.Sleep(time.Second * time.Duration(a.Configuration.RunDuration))
	log.Printf("[%s] app shutting down", a.Name)

	return nil
}
