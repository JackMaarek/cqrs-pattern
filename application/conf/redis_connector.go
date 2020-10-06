package conf

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()
var RedisClient *redis.Client

type RedisConf struct {
	RUrl      string `env:"REDIS_URL"`
	RPassword string `env:"REDIS_PASSWD"`
}

func InitRedisClient() {
	cfg := RedisConf{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RUrl,
		Password: cfg.RPassword, // no password set
		DB:       0,             // use default DB
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Cannot connect to redis client")
	}
	fmt.Println(pong, err)
	log.Println("Connect to redis client")
	RedisClient = client
}
