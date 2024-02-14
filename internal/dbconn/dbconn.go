package dbconn

import (
	"fmt"

	"example.com/playground-wire-app/internal/config"
)

type Conn interface {
	Ping() error
	Close() error
}

type ConnBasic struct {
	conf config.Config
}

func ProvideDbConn(conf config.Config) Conn {
	return &ConnBasic{conf: conf}
}

func (c *ConnBasic) Ping() error {
	if c.conf.SomeInt%2 == 0 {
		return fmt.Errorf("SomeInt is not even: %d", c.conf.SomeInt)
	}

	return nil
}

func (c *ConnBasic) Close() error {
	return nil
}
