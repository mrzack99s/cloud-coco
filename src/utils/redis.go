package utils

import (
	"github.com/go-redis/redis"
	"github.com/mrzack99s/cloud-coco/src/configures"
)

func RedisFindExistingKey(key string) bool {
	_, err := configures.CacheInstance().Get(key).Result()
	if err != redis.Nil {
		return true
	} else {
		return false
	}
}

func RedisDeleteWithPrefix(prefix string) error {
	iter := configures.CacheInstance().Scan(0, prefix, 0).Iterator()
	for iter.Next() {
		val := iter.Val()
		configures.CacheInstance().Del(val)
	}
	if err := iter.Err(); err != nil {
		return err
	}

	return nil
}

func RedisCountWithPrefix(prefix string) (int, error) {
	keys, _, err := configures.CacheInstance().Scan(0, prefix, 0).Result()
	if err != nil {
		return -1, err
	}
	return len(keys), nil
}

func RedisDeleteWithKey(key string) {
	configures.CacheInstance().Del(key)
}
