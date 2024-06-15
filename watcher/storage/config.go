package storage

import (
	"os"

	"encoding/json"

	"github.com/kelpi-io/kelpi-io/watcher/checkers"
)

// Read config from YAML file and return object
func GetConfig(configPath string) (checkers.WatcherConfig, error) {
	yamlFile, errRead := os.ReadFile(configPath)

	if errRead != nil {
		return checkers.WatcherConfig{}, errRead
	}

	var configs checkers.WatcherConfig

	errParse := json.Unmarshal(yamlFile, &configs)

	if errParse != nil {
		return checkers.WatcherConfig{}, errParse
	}

	return configs, nil
}
