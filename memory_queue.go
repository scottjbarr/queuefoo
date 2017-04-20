package queuefoo

import "fmt"

type MemoryQueue struct {
	messages *[]Message
}

func NewMemoryQueue() MemoryQueue {
	return MemoryQueue{
		messages: &[]Message{},
	}
}

func (m MemoryQueue) Receive(messages chan<- Message) ([]Message, error) {
	return *m.messages, nil
}

func (m MemoryQueue) Send(Message) error {
	return fmt.Errorf("Send not implemented")
}

func (m MemoryQueue) Ack(Message) error {
	return fmt.Errorf("Ack not implemented")
}
