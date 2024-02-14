package app

import (
	"errors"
	"fmt"
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

func (a *App) Run() error {
	switch {
	case a.Logger == nil, a.Repository == nil:
		return errors.New("app has no name")
	}

	defer a.Logger.Log("app.App.Run", "app_shutdown")

	a.Logger.Log("app.App.Run", "app_start")
	data, err := a.Repository.Read()
	if err != nil {
		return err
	}

	a.Logger.Log("app.App.Run", fmt.Sprintf("got some data: %v", data))

	time.Sleep(time.Second * time.Duration(a.Configuration.RunDuration))

	return a.Repository.Close()
}
