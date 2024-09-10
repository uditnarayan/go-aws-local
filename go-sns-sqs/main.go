package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/urfave/cli"
	"go-sns-sqs/actions"
	"log"
	"math/rand"
	"os"
)

var (
	app      *cli.App
	ctx      context.Context
	snsActor *actions.SnsActor
	sqsActor *actions.SqsActor
)

func init() {
	app = cli.NewApp()
	app.Name = "go sns example"
	app.Version = "0.0.0"

	ctx = context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	sqsActor = &actions.SqsActor{
		SqsClient: sqs.NewFromConfig(sdkConfig, func(o *sqs.Options) {
			o.BaseEndpoint = aws.String("http://localhost:4566")
		}),
	}
	snsActor = &actions.SnsActor{
		SnsClient: sns.NewFromConfig(sdkConfig, func(o *sns.Options) {
			o.BaseEndpoint = aws.String("http://localhost:4566")
		}),
	}
}

func main() {
	const snsTopicArn = "arn:aws:sns:us-east-1:000000000000:dummy-topic"
	const sqsQueueUrl = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/dummy-queue"
	const maxMessages = 10
	const waitTime = 1

	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "publishes a message to a sns topic",
			Action: func(c *cli.Context) error {
				publishMessage(snsTopicArn, maxMessages)
				return nil
			},
		},
		{
			Name:  "poll",
			Usage: "polls a message from the actions queue",
			Action: func(c *cli.Context) error {
				pollMessages(sqsQueueUrl, maxMessages, waitTime)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func publishMessage(topicArn string, maxMessages int) {
	count := rand.Intn(maxMessages-1) + 1
	events, err := actions.GenerateEvents(count)
	if err != nil {
		panic(err)
	}

	for _, event := range events {
		log.Printf("Event - %+v", event)
		message, err := json.Marshal(event)
		err = snsActor.Publish(ctx, topicArn, string(message))
		if err != nil {
			panic(err)
		}
	}
}

func pollMessages(queueUrl string, maxMessages int32, waitTime int32) {
	sqsActor.Poll(ctx, queueUrl, maxMessages, waitTime)
}
