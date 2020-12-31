package db

import (
	"fmt"
	"gin-mvc/conf"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var RedisDb *redis.Client

func RedisInit() {
	fmt.Println("启动redis连接", time.Now().Format("2006-01-02 15:04:05"))
	//连接服务器
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		DB:       conf.RedisDb,
	})
	//心跳
	pong, err := RedisDb.Ping().Result()
	if err != nil {
		log.Fatalln("redis连接失败", err.Error())
	} else {
		fmt.Println("redis连接成功", pong, time.Now().Format(conf.TimeFormatLayout))
	}
}

func Redis() *redis.Client {
	return RedisDb
}
