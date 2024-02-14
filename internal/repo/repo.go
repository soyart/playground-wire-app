package repo

import (
	"errors"

	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/logger"
)

type Repo interface {
	Close() error
	Read() ([]byte, error)
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
	error,
) {
	if conn == nil {
		return nil, nil, errors.New("nil conn")
	}

	repo := RepoBasic{
		conn:   conn,
		logger: logger,
	}

	return &repo, func() { repo.Close() }, nil
}

func (r *RepoBasic) Read() ([]byte, error) {
	r.logger.Log("repo.RepoBasic.Read", "reading from conn")

	err := r.conn.Ping()
	if err != nil {
		return nil, err
	}

	return []byte("Some data"), nil
}

func (r *RepoBasic) Close() error {
	r.logger.Log("repo.RepoBasic.Close", "closing repo")

	return r.conn.Close()
}
