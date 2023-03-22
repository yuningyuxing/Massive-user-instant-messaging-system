package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// 定义一个全局的链接池
var pool *redis.Pool

// 初始化连接池
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   //表示和数据库的最大链接数
		IdleTimeout: idleTimeout, //最大空闲数
		//初始化链接的代码 表示链接到那个ip的redis
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address,
				redis.DialUsername(""),
				redis.DialPassword("20020902=QWer"))
		},
	}
}
