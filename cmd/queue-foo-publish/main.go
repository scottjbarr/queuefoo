package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/scottjbarr/queuefoo"
)

func main() {
	batchCount, err := strconv.ParseInt(os.Getenv("BATCH_COUNT"), 10, 32)

	if err != nil {
		panic(err)
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// publish a bunch of messages
	config := queuefoo.NewConfig()

	q := queuefoo.NewSQSQueue(config)

	log.Printf("config = %+v\n", config)

	batches := [][]queuefoo.Message{}

	for j := 0; j < int(batchCount); j++ {
		messages := []queuefoo.Message{}

		for i := 0; i < 10; i++ {
			//uid := uuid.NewV4().String()
			id := fmt.Sprintf("%v", j*10+i)
			m := queuefoo.Message{
				ID:      id,
				Payload: id,
			}
			messages = append(messages, m)
		}

		batches = append(batches, messages)
	}

	// fmt.Printf("%v\n", batches)
	// os.Exit(1)

	log.Printf("Sending %v batches of %v messages", batchCount, 10)

	wg := sync.WaitGroup{}
	wg.Add(int(batchCount))

	for i := range batches {
		go func(i int) {
			defer wg.Done()

			messages := batches[i]

			// fmt.Printf("%#+v\n", messages)
			t := time.Now()
			if err := q.SendBatch(messages); err != nil {
				panic(err)
			}

			d := time.Now().Sub(t)
			log.Printf("batch id %v : published = %v : elapsed = %v\n", i,
				len(messages),
				d)
		}(i)
	}

	wg.Wait()
}
