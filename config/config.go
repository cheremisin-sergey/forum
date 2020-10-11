package config

import "github.com/kelseyhightower/envconfig"
import jsoniter "github.com/json-iterator/go"

const (
	envPrefix = "FORUM"
)

type Config struct {
	AppName string `default:forum`
}

func NewConfig() *Config {
	var config Config

	err := envconfig.Process(envPrefix, &config)
	if err == nil {

		// TODO
		_, err := jsoniter.Marshal(config)
		if err != nil {

		}
	}

	return &config
}
