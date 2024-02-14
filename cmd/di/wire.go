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

func InitializeApp() app.App {
	wire.Build(
		config.ProvideConfig,
		dbconn.ProvideDbConn,
		repo.ProvideRepo,
		app.ProvideApp,
	)

	return app.App{}
}
