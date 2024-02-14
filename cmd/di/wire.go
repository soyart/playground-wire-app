//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"example.com/playground-wire-app/internal/app"
	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/logger"
	"example.com/playground-wire-app/internal/repo"
)

// PersistentSet is not buildable on its own, because ProvideRepo requires Logger,
// but none of the providers in this set provides Logger.
//
// To use it, you must build it along with other providers that provide Logger.
var PersistenceSet = wire.NewSet(
	// Since Go best practice is to return concrete types, we'll need
	// to bind the interface to the type that implements it.
	//
	// The first argument to wire.Bind is a pointer to a value of the desired interface type
	// and the second argument is a pointer to a value of the type that implements the interface.

	wire.Bind(new(dbconn.Conn), new(*dbconn.ConnBasic)),
	dbconn.ProvideDbConn, // returns *ConnBasic, which implements Conn

	wire.Bind(new(repo.Repo), new(*repo.RepoBasic)),
	repo.ProvideRepo, // returns *RepoBasic, which implements Repo
)

func InitializeApp() (*app.App, func(), error) {
	wire.Build(
		// config.Config will be provided by this wire.Value
		wire.Value(
			config.Config{
				AppName:     "prod",
				RunDuration: 2,
			},
		),

		PersistenceSet,

		// config.Config.AppName will be the value for "string" provider,
		// which is used in ProvideLogger and app.App.Name
		wire.FieldsOf(new(config.Config), "AppName"),

		// ProvideLogger provide Logger,
		// with appName injected from config.Config.AppName
		wire.Bind(new(logger.Logger), new(*logger.LoggerBasic)),
		logger.ProvideLoggerBasic,

		// Inject fields "Name", "Configuration" and "Repository"
		// with some provider values within this injector
		wire.Struct(new(app.App), "Configuration", "Repository", "Logger"),
	)

	return nil, nil, nil
}

func InitializeAppDebug() (*app.App, func(), error) {
	wire.Build(
		wire.Value("debug"), // strings will be provided with value "debug"
		wire.Bind(new(logger.Logger), new(*logger.LoggerCount)),
		wire.Struct(new(app.App), "Configuration", "Repository", "Logger"),
		PersistenceSet,
		config.ProvideDefaultConfig,
		logger.ProvideLoggerCount,
	)

	return nil, nil, nil
}
