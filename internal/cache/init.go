package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("session.address"),
		Password: viper.GetString("session.password"), // no password set
		DB:       viper.GetInt("session.db"),          // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	Redis = client
}
