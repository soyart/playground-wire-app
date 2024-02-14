package repo

import (
	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/logger"
)

type Repo interface {
	Close() error
}

type RepoBasic struct {
	conn   dbconn.Conn
	logger logger.Logger
}

func ProvideRepo(
	conn dbconn.Conn,
	logger logger.Logger,
) (
	*RepoBasic,
	func(),
) {
	if conn == nil {
		return nil, nil
	}

	repo := RepoBasic{
		conn:   conn,
		logger: logger,
	}

	return &repo, func() { repo.Close() }
}

func (r *RepoBasic) Close() error {
	r.logger.Log("repo.RepoBasic", "closing repo")
	return r.conn.Close()
}
