package repo

import "example.com/playground-wire-app/internal/dbconn"

type Repo struct {
	conn dbconn.Conn
}

func ProvideRepo(conn dbconn.Conn) (Repo, func()) {
	repo := Repo{conn: conn}
	return repo, func() { repo.Close() }
}

func (r *Repo) Close() error {
	return r.conn.Close()
}
