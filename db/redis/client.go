package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

type ConnectionConfigRedis struct {
	Host        string
	Port        string
	DBNumber    int
	ServiceName string
}

var redisClient *redis.Client = nil

func InitRedisClient(config *ConnectionConfigRedis) {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.Host + config.Port,
			Password: "",
			DB:       config.DBNumber,
		})

		pong, err := redisClient.Ping().Result()
		fmt.Println(pong, err)
		if err != nil {
			log.Fatal("Redis server is not reachable from ", config.ServiceName)
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
