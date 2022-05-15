package configures

import (
	"fmt"

	"github.com/go-redis/redis"
)

type RedisCacheConfig struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var redisCache *redis.Client

func CacheInstance() *redis.Client {
	return redisCache
}

func (conf *RedisCacheConfig) SetupCache() {
	redisCache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Hostname, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})
}
