package di

import (
	"example.com/playground-wire-app/internal/app"
	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/repo"
)

func InitializeApp() app.App {
	conf := config.ProvideConfig()
	conn := dbconn.ProvideDbConn(conf)
	repo := repo.ProvideRepo(conn)
	app := app.ProvideApp(conf, repo)

	return app
}
