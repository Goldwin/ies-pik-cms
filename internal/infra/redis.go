package infra

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Hosts          string        `env:"REDIS_URL" yaml:"hosts" default:"127.0.0.1:6379"`
	Username       string        `env:"REDIS_USERNAME" yaml:"username" default:"default"`
	Password       string        `env:"REDIS_PASSWORD" yaml:"password" default:""`
	MaxRetries     int           `env:"REDIS_MAX_RETRIES" yaml:"maxRetries" default:"3"`
	ReadTimeout    time.Duration `env:"REDIS_READ_TIMEOUT" yaml:"readTimeout"`
	WriteTimeout   time.Duration `env:"REDIS_WRITE_TIMEOUT" yaml:"writeTimeout"`
	RouteByLatency bool          `env:"REDIS_ROUTE_BY_LATENCY" yaml:"routeByLatency"`
}

func NewRedisClient(r *RedisConfig) redis.UniversalClient {
	var redisClient redis.UniversalClient
	var option redis.UniversalOptions
	if r == nil {
		return nil
	}
	addresses := strings.Split(r.Hosts, ",")
	option = redis.UniversalOptions{
		Addrs:          addresses,
		Username:       r.Username,
		Password:       r.Password,
		MaxRetries:     r.MaxRetries,
		ReadTimeout:    r.ReadTimeout,
		WriteTimeout:   r.WriteTimeout,
		RouteByLatency: r.RouteByLatency,
	}
	if r != nil {
		redisClient = redis.NewUniversalClient(&option)
	} else {
		log.Fatal("Failed to parse redis config")
	}
	str, err := redisClient.Ping(context.Background()).Result()
	if err != nil || str != "PONG" {
		log.Fatal("Failed to connect redis")
	}
	return redisClient
}
