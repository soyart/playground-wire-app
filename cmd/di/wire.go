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
	dbconn.ProvideDbConn, // returns *ConnBasic, which implements Conn
	// Since Go best practice is to return concrete types, we'll need
	// to bind the interface to the type that implements it.
	//
	// The first argument to wire.Bind is a pointer to a value of the desired interface type
	// and the second argument is a pointer to a value of the type that implements the interface.
	wire.Bind(new(dbconn.Conn), new(*dbconn.ConnBasic)),

	repo.ProvideRepo,
)

func InitializeApp() (*app.App, func()) {
	wire.Build(
		// config.Config will be provided by this wire.Value
		wire.Value(
			config.Config{
				AppName:     "production",
				RunDuration: 2,
			},
		),

		PersistenceSet,

		// config.Config.AppName will be the value for "string" provider,
		// which is used in app.App.Name
		wire.FieldsOf(new(config.Config), "AppName"),

		// Inject fields "Name", "Configuration" and "Repository"
		// with some provider values within this injector
		wire.Struct(new(app.App), "Name", "Configuration", "Repository"),
	)

	return nil, nil
}
