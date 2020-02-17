package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

const DefaultPort = 9129
const DefaultCacheDir = "/var/cache/pacoloco"

type Repo struct {
	Url string `yaml:"url"`
}

type Config struct {
	CacheDir string          `yaml:"cache_dir"`
	Port     int             `yaml:"port"`
	Repos    map[string]Repo `yaml:"repos,omitempty"`
}

var config *Config

func readConfig(filename string) *Config {
	var result = &Config{CacheDir: DefaultCacheDir, Port: DefaultPort}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
