package models

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"online-chat-server/config"
)

var rdb *redis.Client

func GetRedis() *redis.Client {
	return rdb
}

func init()  {
	r := config.Conf.Redis

	rdb = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", r.Host, r.Port),
		Password:           r.PassWord,
		DB:                 0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Printf("init redis failed: %v \n", err)
		return
	}

	log.Printf("redis: %s:%s \n", r.Host, r.Port)

}