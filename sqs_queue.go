package queuefoo

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSQueue struct {
	Config Config
}

func NewSQSQueue(config Config) SQSQueue {
	return SQSQueue{
		Config: config,
	}
}

func (s SQSQueue) Receive(messages chan<- Message) error {
	maxMessages := int64(10)
	waitTime := int64(20)

	rmin := &sqs.ReceiveMessageInput{
		QueueUrl:            &s.Config.QueueURL,
		MaxNumberOfMessages: &maxMessages,
		WaitTimeSeconds:     &waitTime,
	}

	// client := sqs.New(session.New(), s.Config.AWSConfig())

	// messages := []Message{}

	// loop as long as there are messages on the queue
	for {
		resp, err := s.client().ReceiveMessage(rmin)

		if err != nil {
			return err
		}

		if len(resp.Messages) == 0 {
			log.Printf("No messages")
			return nil
		}

		log.Printf("received %v messages...", len(resp.Messages))

		for _, m := range resp.Messages {
			message := Message{}
			message.ID = *m.MessageId
			message.Handle = *m.ReceiptHandle
			message.Payload = *m.Body

			// messages = append(messages, message)
			messages <- message
		}
	}

	return nil
}

func (s SQSQueue) Send(m Message) error {
	return fmt.Errorf("Send not implemented")
}

func (s SQSQueue) SendBatch(messages []Message) error {
	entries := []*sqs.SendMessageBatchRequestEntry{}

	for i := 0; i < len(messages); i++ {
		m := messages[i]

		entry := sqs.SendMessageBatchRequestEntry{
			Id:          &m.ID,
			MessageBody: &m.Payload,
		}
		entries = append(entries, &entry)
	}

	smbi := sqs.SendMessageBatchInput{
		QueueUrl: &s.Config.QueueURL,
		Entries:  entries,
	}

	req, output := s.client().SendMessageBatchRequest(&smbi)

	if err := req.Send(); err != nil {
		return err
	}

	if len(output.Successful) != len(messages) {
		return fmt.Errorf("Messages fail count : %v", len(output.Failed))
	}

	return nil
}

func (s SQSQueue) Ack(m Message) error {
	dmi := &sqs.DeleteMessageInput{
		QueueUrl:      &s.Config.QueueURL,
		ReceiptHandle: &m.Handle,
	}

	// client := sqs.New(session.New(), s.Config.AWSConfig())

	_, err := s.client().DeleteMessage(dmi)

	return err
}

func (s SQSQueue) client() *sqs.SQS {
	return sqs.New(session.New(), s.Config.AWSConfig())
}
