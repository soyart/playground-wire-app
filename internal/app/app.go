package app

import (
	"errors"
	"fmt"
	"time"

	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/logger"
	"example.com/playground-wire-app/internal/repo"
)

const (
	errNoDeps = "missing dependencies"
	errNoName = "app has no name"
	errNoData = "length of data bytes is 0"
)

type App struct {
	Configuration config.Config
	Repository    repo.Repo
	Logger        logger.Logger
}

func (a *App) Run() error {
	switch {
	case a.Logger == nil, a.Repository == nil:
		return errors.New(errNoDeps)
	}

	if a.Configuration.AppName == "" {
		return errors.New(errNoName)
	}

	defer a.Logger.Log("app.App.Run", "app_shutdown")

	a.Logger.Log("app.App.Run", "app_start")

	data, err := a.Repository.Read(a.Configuration.AppName)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return errors.New(errNoData)
	}

	a.Logger.Log("app.App.Run", fmt.Sprintf("got some data: %v", data))

	time.Sleep(time.Second * time.Duration(a.Configuration.RunDuration))

	return a.Repository.Close()
}
