package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MaxParallelDownloads   int `yaml:"maxParallelDownloads"`
	MaxUrlsInSearchResults int `yaml:"maxUrlsInSearchResults"`
	QueueSize              int `yaml:"queueSize"`
	HttpPort               int `yaml:"httpPort"`
}

func ReadConfig(file string) (*Config, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var c Config

	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	return &c, nil
}
