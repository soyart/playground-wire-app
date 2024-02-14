package repo

import "example.com/playground-wire-app/internal/dbconn"

type Repo struct {
	conn dbconn.Conn
}

func ProvideRepo(conn dbconn.Conn) Repo {
	return Repo{conn: conn}
}

func (r *Repo) Close() error {
	return r.conn.Close()
}
