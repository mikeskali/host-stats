package cfg

type config struct {
	modules     []string `env:"modules" envDefault:"dev"`
	serviceName string   `env:"serviceName" envDefault:"host-stats"`
}