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

	nc.Subscribe("intros", func(m *nats.Msg) {
		fmt.Printf("I got message: %s\n", string(m.Data))
	})

	time.Sleep(1 * time.Hour)
}
