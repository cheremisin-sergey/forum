package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)
import jsoniter "github.com/json-iterator/go"

const (
	envPrefix = "FORUM"
)

//type Config struct {
//	Address          string `env:"ADDRESS" envDefault:":8080"`
//	AppName string `default:forum`
//}

type Config struct {
	ServerPort       string        `yaml:"server_port" envconfig:"FORUM_SERVER_PORT"`
	ServerHost       string        `yaml:"server_host" envconfig:"FORUM_SERVER_HOST"`
	DatabaseUsername string        `yaml:"username" envconfig:"FORUM_DB_USERNAME"`
	DatabasePassword string        `yaml:"password" envconfig:"FORUM_DB_PASSWORD"`
	ReadTimeout      time.Duration `yaml:"read_timeout" envconfig:"FORUM_SERVER_READ_TIMEOUT"`
	WriteTimeout     time.Duration `yaml:"write_timeout" envconfig:"FORUM_SERVER_WRITE_TIMEOUT"`
}

func NewConfig() *Config {
	var config Config

	readFile(&config)
	readEnv(&config)
	return &config
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(config *Config) {
	var f, err = os.Open(".config/config_local.yaml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		processError(err)
	}
}

func readEnv(config *Config) {
	err := envconfig.Process(envPrefix, config)
	if err == nil {

		// TODO
		data, err := jsoniter.Marshal(config)
		if err != nil {
			fmt.Println("Config fail")
			processError(err)
		} else {
			fmt.Println(fmt.Sprintf("Config success %v", data));
		}
	}
	fmt.Println(config.ServerPort)
}
