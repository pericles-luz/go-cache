package cache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/pericles-luz/go-base/pkg/conf"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(file string) *Redis {
	config := conf.NewRedis()
	err := config.Load(file)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	return &Redis{client: client}
}

func (r *Redis) Get(key string) (string, error) {
	val, err := r.client.Get(key).Result()
	if err != nil && err.Error() == "redis: nil" {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *Redis) Set(key string, value string, durationInMinutes int) error {
	err := r.client.Set(key, value, time.Minute*time.Duration(durationInMinutes)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Del(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Ping() error {
	_, err := r.client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
