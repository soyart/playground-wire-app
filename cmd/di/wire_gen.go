// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"example.com/playground-wire-app/internal/app"
	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/logger"
	"example.com/playground-wire-app/internal/repo"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeApp() (*app.App, func(), error) {
	config := _wireConfigValue
	connBasic := dbconn.ProvideDbConn(config)
	string2 := config.AppName
	loggerBasic := logger.ProvideLoggerBasic(string2)
	repoBasic, cleanup, err := repo.ProvideRepo(connBasic, loggerBasic)
	if err != nil {
		return nil, nil, err
	}
	appApp := &app.App{
		Configuration: config,
		Repository:    repoBasic,
		Logger:        loggerBasic,
	}
	return appApp, func() {
		cleanup()
	}, nil
}

var (
	_wireConfigValue = config.Config{
		AppName:     "prod",
		RunDuration: 2,
	}
)

func InitializeAppDebug() (*app.App, func(), error) {
	string2 := _wireStringValue
	configConfig := config.ProvideDefaultConfig(string2)
	connBasic := dbconn.ProvideDbConn(configConfig)
	loggerCount := logger.ProvideLoggerCount(string2)
	repoBasic, cleanup, err := repo.ProvideRepo(connBasic, loggerCount)
	if err != nil {
		return nil, nil, err
	}
	appApp := &app.App{
		Configuration: configConfig,
		Repository:    repoBasic,
		Logger:        loggerCount,
	}
	return appApp, func() {
		cleanup()
	}, nil
}

var (
	_wireStringValue = "debug"
)

// wire.go:

// PersistentSet is not buildable on its own, because ProvideRepo requires Logger,
// but none of the providers in this set provides Logger.
//
// To use it, you must build it along with other providers that provide Logger.
var PersistenceSet = wire.NewSet(wire.Bind(new(dbconn.Conn), new(*dbconn.ConnBasic)), dbconn.ProvideDbConn, wire.Bind(new(repo.Repo), new(*repo.RepoBasic)), repo.ProvideRepo)
