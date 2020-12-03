package common

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type AppConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}


type Config struct {
	App1  AppConfig `yaml:"app1"`
	App2 AppConfig `yaml:"app2"`
	UserClient AppConfig `yaml:"client"`
}


func NewConfig(configPath string) *Config {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil{
		log.Fatalf("Failed to decode config.")
	}

	return &cfg
}