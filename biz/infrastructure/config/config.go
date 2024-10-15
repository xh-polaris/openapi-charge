package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"os"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"
)

var config *Config

type Security struct {
	Issue string
}

type Config struct {
	service.ServiceConf
	ListenOn string
	Mongo    struct {
		URL string
		DB  string
	}
	Cache    cache.CacheConf
	Security Security
}

func NewConfig() (*Config, error) {
	c := new(Config)
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "etc/config.yaml"
	}
	err := conf.Load(path, c)
	if err != nil {
		return nil, err
	}
	err = c.SetUp()
	if err != nil {
		return nil, err
	}
	config = c
	return c, nil
}

func GetConfig() *Config {
	return config
}
