package main

import (
	"log"

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
			// log.Printf("%v", m.ID)
			if err := q.Ack(m); err != nil {
				log.Printf("[ERROR] failed to ack message err=%v message=%v",
					err,
					m)
			}
		}
	}()

	if err := q.Receive(messages); err != nil {
		panic(err)
	}
}
