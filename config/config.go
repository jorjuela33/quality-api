package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	PrivateKeyPath string
	PublicKeyPath  string
}

var environments = map[string]string{
	"production":  "config/production.json",
	"development": "config/development.json",
	"tests":       "../../config/tests.json",
}

var config Config = Config{}
var env = "development"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "development"
	}

	loadSettingsByEnv(env)
}

func Current() Config {
	if &config == nil {
		Init()
	}

	return config
}

func loadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}

	config = Config{}
	jsonErr := json.Unmarshal(content, &config)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}
