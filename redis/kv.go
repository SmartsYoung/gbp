package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func initRedisPool() {
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisAddress)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", RedisDb)
			return c, nil
		},
	}
}

/**
 * 设置redis的对应key的value
 */
func redisSet(key string, value string) {
	c, err := RedisClient.Dial()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	_, err = c.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

/**
 * 获取redis的对应key的value
 */
func redisGet(key string) (value string) {
	c, err := RedisClient.Dial()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	val, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return ""
	} else {
		fmt.Printf("Got value is %v \n", val)
		return val
	}
}

/**
 * redis使得对应的key的值自增
 */
func redisIncr(key string) (value string) {
	c, err := RedisClient.Dial()
	_, err = c.Do("INCR", key)
	if err != nil {
		fmt.Println("incr error", err.Error())
	}

	incr, err := redis.String(c.Do("GET", key))
	if err == nil {
		fmt.Println("redis key after incr is : ", incr)
	}
	return incr
}
