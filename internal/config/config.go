package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP     `yaml:"http"`
	DataBase `yaml:"db"`
	GRPC     `yaml:"grpc"`
}

type HTTP struct {
	Port string `env-required:"true" yaml:"port"`
	Host string `env-required:"true" yaml:"host"`
}

type DataBase struct {
	Host     string `env-required:"true" yaml:"host"`
	Port     int    `env-required:"true" yaml:"port"`
	User     string `env-required:"true" yaml:"user"`
	Password string `env-required:"true" yaml:"password"`
	DB       string `env-required:"true" yaml:"db"`
	SSLMode  string `env-required:"true" yaml:"ssl_mode"`
}

type GRPC struct {
	Port string `env-required:"true" yaml:"port"`
	Host string `env-required:"true" yaml:"host"`
}

func Init() (*Config, error) {
	config := &Config{}
	if err := cleanenv.ReadConfig("./config.yaml", config); err != nil {
		return nil, err
	}
	return config, nil
}
