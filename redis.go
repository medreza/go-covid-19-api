package main

import (
	"github.com/gomodule/redigo/redis"
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   100,
		MaxActive: 10000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func setRedisCache(blobKey string, blob string, secExpiration int, c redis.Conn) error {
	_, err := c.Do("SET", blobKey, blob, "EX", secExpiration)
	if err != nil {
		return err
	}
	return nil
}

func getRedisCache(blobKey string, c redis.Conn) (string, error) {
	data, err := redis.String(c.Do("GET", blobKey))
	if err != nil {
		return "", err
	}
	return data, nil
}
