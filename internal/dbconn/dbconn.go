package dbconn

import (
	"errors"
	"time"

	"example.com/playground-wire-app/internal/config"
)

type Conn interface {
	Ping() error
	Close() error
}

type ConnBasic struct {
	conf config.Config
}

func ProvideDbConn(conf config.Config) *ConnBasic {
	return &ConnBasic{conf: conf}
}

func (c *ConnBasic) Ping() error {
	if time.Now().Unix()%2 == 0 {
		return errors.New("failed to ping db")
	}

	return nil
}

func (c *ConnBasic) Close() error {
	return nil
}
