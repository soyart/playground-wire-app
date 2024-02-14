package config

type Config struct {
	AppName     string
	RunDuration int64
}

func ProvideDefaultConfig() Config {
	return Config{
		AppName:     "dev",
		RunDuration: 0,
	}
}
