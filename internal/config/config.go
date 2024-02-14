package config

import "time"

type Config struct {
	SomeInt int64
}

func ProvideConfig() Config {
	return Config{
		SomeInt: time.Now().Unix(),
	}
}
