package main

import (
	"log"
	"math/rand"

	"github.com/scottjbarr/queuefoo"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	config := queuefoo.NewConfig()

	log.Printf("config = %+v\n", config)

	q := queuefoo.NewSQSQueue(config)

	messages := make(chan queuefoo.Message, 1)

	go func() {
		for {
			m := <-messages

			// fail randomly
			if rand.Intn(4) == 0 {
				log.Printf("[ERROR] monkeys id=%v", m.ID)
				continue
			}

			if err := q.Ack(m); err != nil {
				log.Printf("[ERROR] ACK fail id=%v : err = %v", m.ID, err)
				continue
			}

			log.Printf("[INFO] ACK OK %v", m.ID)
		}
	}()

	if err := q.Receive(messages); err != nil {
		panic(err)
	}
}
