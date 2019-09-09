package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"starter-kit/utils"
)

var config Config
var envConfig map[string]interface{}
var defaultConfig map[string]interface{}

type Config struct {
	runEnv string
}

func (c Config) Get(key string) interface{} {
	envValue := os.Getenv(envConfig[key].(string))
	if len(envValue) != 0 {
		return envValue
	}
	switch c.runEnv {
	default:
		return defaultConfig[key]
	}
}

func init() {
	envTxt, _ := ioutil.ReadFile(`config/env.json`)
	defaultTxt, _ := ioutil.ReadFile(`config/default.json`)
	_ = json.Unmarshal(envTxt, &envConfig)
	_ = json.Unmarshal(defaultTxt, &defaultConfig)
	setConfig(Config{
		runEnv: utils.GetEnv(`GO_ENV`, `development`),
	})
}

func setConfig(c Config) {
	config = c
}

func GetConfig() Config {
	return config
}
