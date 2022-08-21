package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	YoutubeAPIKey string `yaml:"youtubeAPIKey"`
}

func NewConfig(path string) Config {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var c Config
	err = yaml.NewDecoder(f).Decode(&c)
	return c
}
