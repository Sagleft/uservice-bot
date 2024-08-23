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
	Utopia utopiago.Config
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
