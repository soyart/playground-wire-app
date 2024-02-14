package config

type Config struct {
	AppName     string
	RunDuration int64
}

func ProvideDefaultConfig(appName string) Config {
	return Config{
		AppName:     appName,
		RunDuration: 0,
	}
}
