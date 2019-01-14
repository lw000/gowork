// redis_test project main.go
package main

import (
	"fmt"
	// "fmt"
	"log"
	// "os"

	redigo "github.com/garyburd/redigo/redis"

	"redis_test/redisClient"
)

func Test() {
	c, err := redigo.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		log.Printf("Connect to redis error", err)
		return
	}
	defer c.Close()

	var reply interface{}

	reply, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		log.Println("redis set failed:", err)
	}

	if reply != nil {
		log.Println(reply)
	}

	username, err := redigo.String(c.Do("GET", "mykey"))
	if err != nil {
		log.Println("redis get failed:", err)
	} else {
		log.Printf("get mykey: %v", username)
	}

	n, _ := c.Do("EXPIRE", "expire_levi", 24*3600)
	if n == int64(1) {
		log.Println("success")
	}
}

func main() {
	c := redisClient.Get()
	defer func() {
		c.Close()
	}()

	reply, err := c.Do("SET", "go_key", "redigo")
	if err != nil {
		log.Println(err)
	}

	if reply != nil {
		fmt.Println(reply)
	}

	s, err := redigo.String(c.Do("GET", "go_key"))
	if err != nil {

	}

	fmt.Println(s)
}
