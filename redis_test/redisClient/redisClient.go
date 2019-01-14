package redisClient

import (
	"fmt"
	// "time"

	redigo "github.com/garyburd/redigo/redis"
)

var pool *redigo.Pool

func init() {
	redisHost := /*"127.0.0.1"*/ "192.168.1.201"
	redisPort := 6379
	pool_size := 20
	password := /*"123456"*/ ""

	pool = redigo.NewPool(func() (redigo.Conn, error) {
		dns := fmt.Sprintf("%s:%d", redisHost, redisPort)
		c, err := redigo.Dial("tcp", dns, redigo.DialDatabase(1), redigo.DialPassword(password))
		if err != nil {
			return nil, err
		}

		if len(password) > 0 {
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
		}

		return c, err

	}, pool_size)

	// pool = &redigo.Pool{
	// 	Dial: func() (redigo.Conn, error) {
	// 		dns := fmt.Sprintf("%s:%d", redisHost, redisPort)
	// 		c, err := redigo.Dial("tcp", dns, redigo.DialDatabase(0), redigo.DialPassword(password))
	// 		if err != nil {
	// 			return nil, err
	// 		}

	// 		// if _, err := c.Do("AUTH", password); err != nil {
	// 		// 	c.Close()
	// 		// 	return nil, err
	// 		// }

	// 		return c, err
	// 	},
	// 	IdleTimeout: time.Second * time.Duration(240),
	// 	MaxIdle:     pool_size,
	// }
}

func Get() redigo.Conn {
	return pool.Get()
}
