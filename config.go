package queuefoo

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
)

type Config struct {
	QueueURL string
}

func NewConfig() Config {
	return Config{
		QueueURL: os.Getenv("QUEUE_URL"),
	}
}

func (c Config) AWSConfig() *aws.Config {
	return &aws.Config{}
}
