package config

import (
	"os"

	"github.com/go-faster/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB struct {
		Ivalue   string `yaml:"ivalue"`
		Migrator string `yaml:"migrator"`
		MsSql    string `yaml:"mssql"`
	} `yaml:"db"`
}

func New(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "config")
	}
	c := Config{}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, errors.Wrap(err, "config")
	}

	return &c, nil
}
