package dbconn

import (
	"fmt"

	"example.com/playground-wire-app/internal/config"
)

type Conn struct {
	conf config.Config
}

func ProvideDbConn(conf config.Config) Conn {
	return Conn{conf: conf}
}

func (c *Conn) Ping() error {
	if c.conf.SomeInt%2 == 0 {
		return fmt.Errorf("SomeInt is not even: %d", c.conf.SomeInt)
	}

	return nil
}

func (c *Conn) Close() error {
	return nil
}
