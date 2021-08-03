package redis

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr     string
	Username string
	Password string
	Database int
}

func Connect(config Config) (*Client, error) {
	opts := redis.Options{
		Addr:     config.Addr,
		Username: config.Username,
		Password: config.Password,
		DB:       config.Database,
	}

	var counter int
	for {
		log.Println("[INFO] attempt to connect to redis server...")
		client := redis.NewClient(&opts)
		err := client.Ping(context.Background())
		if err == nil {
			return &Client{client}, nil
		}

		if counter == 30 {
			log.Println("[ERROR] reached maximum number of attempt connecting to redis server")
			return nil, errors.New("redis connection failed")
		}

		log.Printf("[WARN] attempt connecting to redis server failed, will be repeated in one second")
		time.Sleep(time.Second)
		counter++
	}
}
