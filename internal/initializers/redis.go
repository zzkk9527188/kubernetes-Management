package initializers

import (
	"cm_platform/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func InitRedis(cfg *config.Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Standalone.Host, cfg.Redis.Standalone.Port),
		Password: cfg.Redis.Standalone.Password,
		DB:       0,
	})
	_, err := redisClient.Ping(redisClient.Context()).Result()
	if err != nil {
		fmt.Errorf("redis 连接失败")
	}
	log.Println("Connected to Redis successfully")
	return redisClient
}
