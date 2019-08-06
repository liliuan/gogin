package lib

import (
	"github.com/gomodule/redigo/redis"
	"gogin/config"
	"time"
)

const (
	redisMaxIdle        = 3   //最大空闲连接数
	redisIdleTimeoutSec = 240 //最大空闲连接时间
	MaxActive           = 800 // 最大连接数
	timeout             = 60
)

var redisClient *redis.Pool

func InitRedis() {
	cfg := config.Cfg
	redisClient = &redis.Pool{
		MaxActive:   MaxActive,
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", cfg.RedisUri,
				redis.DialConnectTimeout(timeout*time.Second),
				redis.DialReadTimeout(timeout*time.Second),
				redis.DialWriteTimeout(timeout*time.Second))

			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
}

//func init(){
//	InitRedis()
//}
