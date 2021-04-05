package config

func NewAppConfig() *AppConfig {
	return &AppConfig{
			/*
		|--------------------------------------------------------------------------
		| Environment Name
		|--------------------------------------------------------------------------
		|
		| With the help of the environment name we can implement different behaviour for
		| our environments. For example for local development we like to have another
		| logger encoding with colors, but in production we have the JSON encoding.
		|
		*/

		Env: Env("APP_ENV", "production"),
		/*
		|--------------------------------------------------------------------------
		| Application Port
		|--------------------------------------------------------------------------
		|
		| This value define on witch port the application is available. Default is
		| the standard port 8080
		|
		*/

		Port: EnvInt("APP_PORT", 3345),
		Prefix: "/index/v1",
		ShowBanner: EnvBool("APP_SHOW_BANNER", true),
		/*
		|--------------------------------------------------------------------------
		| Simultaneous Connections
		|--------------------------------------------------------------------------
		|
		| By default, http.ListenAndServe (which gin.Run wraps) will serve an unbounded
		| number of requests. Limiting the number of simultaneous connections can
		| sometimes greatly speed things up under load.
		|
		*/
		Connection: EnvInt("APP_CONNECTIONS", 20),
	}
}

type AppConfig struct {
	Env        string
	Prefix string
	Port int
	Connection int
	ShowBanner bool
}