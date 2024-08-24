package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id      int64
	Message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1)

			time.Sleep(time.Second * 1)
			msg := Message{id: i, Message: "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{id: i, Message: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case x := <-c1:
			fmt.Printf("c1 received id %d - %+v\n", x.id, x.Message)
		case x := <-c2:
			fmt.Printf("c2 received id %d - %+v\n", x.id, x.Message)
		case <-time.After(time.Second * 3):
			println("timeout")
		}

	}
}
