package main

import (
	"log"

	"example.com/playground-wire-app/internal/app"
	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/repo"
)

func main() {
	conf := config.ProvideConfig()
	conn := dbconn.ProvideDbConn(conf)
	repo := repo.ProvideRepo(conn)
	app := app.ProvideApp(conf, repo)

	err := app.Start()
	if err != nil {
		log.Fatal("app got error", err)
	}
}
