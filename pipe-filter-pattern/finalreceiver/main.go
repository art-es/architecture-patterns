package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	. "github.com/art-es/architecture-patterns/pipe-filter-pattern/common"
)

func main() {
	InitLogger("FINAL RECEIVER")
	ctx := context.Background()
	pubsub := RDB.Subscribe(ctx, ChannelResult)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	log.Println("[INFO] waiting for messages ...")
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("[ERROR] pubsub.ReceiveMessage: %v\n", err)
			continue
		}

		// if you want it to work as a queue, just run without a goroutine
		go handle(msg.Payload)

		select {
		case <-stop:
			pubsub.Close()
			return
		default:
		}
	}
}

func handle(payload string) {
	msg, err := UnmarshalMessage(payload)
	if err != nil {
		log.Printf("[ERROR] UnmarshalMessage: %v\n", err)
		return
	}

	log.Printf("[INFO] received message: %+v\n", *msg)
}
