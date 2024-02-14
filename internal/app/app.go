package app

import (
	"errors"
	"time"

	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/logger"
	"example.com/playground-wire-app/internal/repo"
)

type App struct {
	Configuration config.Config
	Repository    repo.Repo
	Logger        logger.Logger
}

func (a *App) Start() error {
	switch {
	case a.Logger == nil, a.Repository == nil:
		return errors.New("app has no name")
	}

	a.Logger.Log("app.App", "app_start")
	time.Sleep(time.Second * time.Duration(a.Configuration.RunDuration))
	a.Logger.Log("app.App", "app_shutdown")

	return nil
}
