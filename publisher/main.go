package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("cant connect to NATS: %v", err)
	}

	defer nc.Close()

	count := 0
	for {
		nc.Publish("intros", []byte("Hello World!"))
		count++
		fmt.Println("Sent message ", count)
		time.Sleep(1 * time.Second)
	}

}
