package server

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: ":9090",
	}
}
