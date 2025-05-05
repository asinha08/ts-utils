package redis

import (
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

type ConnectionConfigRedis struct {
	Host        string
	Port        string
	DBNumber    int
	ServiceName string
	Username    string
	Password    string
}

var redisClient *redis.Client = nil

func InitRedisClient(config *ConnectionConfigRedis) {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.Host + config.Port,
			Username: config.Username,
			Password: config.Password,
			DB:       config.DBNumber,
		})
		ctx := context.Background()
		pong, err := redisClient.Ping(ctx).Result()
		fmt.Println(pong, err)
		if err != nil {
			panic("Redis server is not reachable from " + config.ServiceName)
		}
		fmt.Println("Redis Client created from ", config.ServiceName)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedisClient() {
	_ = redisClient.Close()
}
