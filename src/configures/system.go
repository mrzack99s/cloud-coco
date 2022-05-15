package configures

import (
	"gopkg.in/yaml.v3"
)

var sysConf SystemConf

func Sys() *SystemConf {
	return &sysConf
}

type SystemConf struct {
	COCO struct {
		DB         DatabaseConfig   `yaml:"db"`
		RedisCache RedisCacheConfig `yaml:"redis_cache"`
		Secret     string           `yaml:"secret"`
	} `yaml:"coco"`
}

func Parse(file []byte) {
	err := yaml.Unmarshal(file, &sysConf)
	if err != nil {
		panic(err)
	}
}
