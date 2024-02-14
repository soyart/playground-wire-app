// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"example.com/playground-wire-app/internal/app"
	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/repo"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeApp() app.App {
	configConfig := config.ProvideConfig()
	connBasic := dbconn.ProvideDbConn(configConfig)
	repoRepo := repo.ProvideRepo(connBasic)
	appApp := app.App{
		Configuration: configConfig,
		Repository:    repoRepo,
	}
	return appApp
}

// wire.go:

var PersistenceSet = wire.NewSet(dbconn.ProvideDbConn, wire.Bind(new(dbconn.Conn), new(*dbconn.ConnBasic)), repo.ProvideRepo)
