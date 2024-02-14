//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"example.com/playground-wire-app/internal/app"
	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/repo"
)

var PersistenceSet = wire.NewSet(
	dbconn.ProvideDbConn,
	repo.ProvideRepo,
)

func InitializeApp() app.App {
	wire.Build(
		config.ProvideConfig,
		PersistenceSet,
		app.ProvideApp,
	)

	return app.App{}
}
