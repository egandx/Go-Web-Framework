package src

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisDefaultPool *redis.Pool

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) {
			c,err:= redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", 123456); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}

func init()  {
	RedisDefaultPool = newPool("127.0.0.1:6379")
}


