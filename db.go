package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Client is exported Mongo Database client
var Client *redis.Client

// ConnectDatabase is used to connect the MongoDB database
func ConnectDatabase() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping().Result()

	if err != nil {
		fmt.Println(Client, err)
	}
}
