package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	HTTPAddr string `yaml:"http_addr"`
	MongoURI string `yaml:"mongodb_uri"`
	MongoDB  string `yaml:"mongodb_db"`
}

func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
