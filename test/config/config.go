package config



type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

//var ErrorNoRows = errors.New("errornorows")

func Load() Config {
	cfg := Config{}
	cfg.PostgresHost = "LocalHost"
	cfg.PostgresPort = "5432"
	cfg.PostgresUser = "davlat"
	cfg.PostgresPassword = "1"
	cfg.PostgresDB = "test1"

	return cfg
}
