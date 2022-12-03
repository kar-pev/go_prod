package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
}

func NewConfig() *Config {
	return &Config{}
}
