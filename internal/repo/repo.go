package repo

import (
	"errors"
	"fmt"

	"example.com/playground-wire-app/internal/dbconn"
	"example.com/playground-wire-app/internal/logger"
)

const errNullKey = "null key"

//go:generate mockgen -source=./repo.go -destination=./mock_repo/mock_repo.go -package=mock_repo
type Repo interface {
	Close() error
	Read(key string) ([]byte, error)
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

func (r *RepoBasic) Read(key string) ([]byte, error) {
	r.logger.Log("repo.RepoBasic.Read", "reading from conn")

	if key == "" {
		return nil, errors.New(errNullKey)
	}

	err := r.conn.Ping()
	if err != nil {
		return nil, err
	}

	data := fmt.Sprintf("Some data for key %s", key)

	return []byte(data), nil
}

func (r *RepoBasic) Close() error {
	r.logger.Log("repo.RepoBasic.Close", "closing repo")

	return r.conn.Close()
}
