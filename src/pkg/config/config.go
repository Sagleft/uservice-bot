package config

import (
	"fmt"

	swissknife "github.com/Sagleft/swiss-knife"
	utopiago "github.com/Sagleft/utopialib-go/v2"
)

const (
	ConfigFilePath = "data/config.yml"
)

type Config struct {
	Utopia utopiago.Config `yaml:"utopia"`
	DB     DBConfig        `yaml:"db"`
}

type DBConfig struct {
	Host          string `yaml:"host"` // default:"localhost"
	Port          int    `yaml:"port"` // default:"3306"
	Name          string `yaml:"dbName"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	TablePrefix   string `yaml:"tablePrefix"`
	ConnTimeoutMS int    `yaml:"connTimeout"` // default:"5000"
	GormDebugMode bool   `yaml:"debugMode"`
	Location      string `yaml:"timeZone"` // default:"Europe/Moscow"
}

func Load() (Config, error) {
	if !swissknife.IsFileExists(ConfigFilePath) {
		return Config{},
			fmt.Errorf(
				"file %q not exists. Read the setup instructions in the README file",
				ConfigFilePath,
			)
	}

	var cfg Config
	if err := swissknife.ParseStructFromYamlFile(ConfigFilePath, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse: %w", err)
	}
	return cfg, nil
}
