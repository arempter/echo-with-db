package config

type Server struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port int    `env:"PORT" envDefault:"1234"`
}
