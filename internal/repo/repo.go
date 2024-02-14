package repo

import "example.com/playground-wire-app/internal/dbconn"

type Repo interface {
	Close() error
}

type RepoBasic struct {
	conn dbconn.Conn
}

func ProvideRepo(conn dbconn.Conn) (*RepoBasic, func()) {
	repo := RepoBasic{conn: conn}
	return &repo, func() { repo.Close() }
}

func (r *RepoBasic) Close() error {
	return r.conn.Close()
}
