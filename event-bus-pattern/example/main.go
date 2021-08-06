package main

import (
	"fmt"

	"github.com/art-es/architecture-patterns/event-bus-pattern/eventbus"
)

func main() {
	bus := eventbus.New()
	sub := bus.Subscribe("example-channel", 1)

	bus.Publish("example-channel", "example-source")
	src, ok := sub.Receive()
	fmt.Printf("Received: %v %t\n", src, ok)

	sub.Unsubscribe()
	src, ok = sub.Receive()
	fmt.Printf("Received: %v %t\n", src, ok)
}
