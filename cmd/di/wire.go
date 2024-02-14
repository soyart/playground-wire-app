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

func InitializeApp() app.App {
	wire.Build(
		config.ProvideConfig,
		PersistenceSet,
		app.ProvideApp,
	)

	return app.App{}
}
