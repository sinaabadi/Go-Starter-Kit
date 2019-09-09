package config

var config Config

type Config struct {
	keys map[string]interface{}
}

func (c *Config) get(key string) interface{} {

}

func init() {

}

func setConfig(c Config) {
	config = c
}

func GetConfig() Config {
	return config
}
