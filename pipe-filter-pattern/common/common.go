package common

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func InitLogger(serviceName string) {
	// time.Sleep(time.Second * 3)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix(fmt.Sprintf("[%s] ", serviceName))
}

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{Addr: env("REDIS_ADDR", "127.0.0.1:6379")})
}

const (
	ChannelA      = "channel-A"
	ChannelB      = "channel-B"
	ChannelResult = "channel-result"
)

type Message struct {
	ID   string
	Data string
}

func UnmarshalMessage(s string) (*Message, error) {
	var msg Message
	if err := json.Unmarshal([]byte(s), &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func NewMessage(data string) (string, string) {
	id := uuid.NewString()
	msg, _ := json.Marshal(&Message{
		ID:   id,
		Data: data,
	})
	return id, string(msg)
}

func PublishMessage(channel, data string) {
	id, msg := NewMessage(data)
	err := RDB.Publish(context.Background(), channel, msg).Err()
	if err != nil {
		log.Printf("[ERROR] sending %s message failed, %v\n", id, err)
		return
	}
	log.Printf("[INFO] message %s has been sent\n", id)
}

func env(varname string, defaultvalue string) string {
	if v := os.Getenv(varname); v != "" {
		return v
	}
	return defaultvalue
}
