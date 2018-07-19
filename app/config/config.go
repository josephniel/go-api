package config

// Config is the structure to hold the parsed configuration data
type Config struct {
	App struct {
		Port int `yaml:"port"`
	}
	Auth struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
	DB struct {
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}
