package config

type AppConfig struct {
	Env        string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Env: Env("APP_ENV", "production"),
	}
}