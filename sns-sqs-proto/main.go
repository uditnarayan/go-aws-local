package main

import (
	"context"
	"encoding/base64"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/urfave/cli"
	"go-sns-sqs/actions"
	"go-sns-sqs/actions/event"
	"google.golang.org/protobuf/proto"
	"log"
	"math/rand"
	"os"
)

var (
	app *cli.App
	ctx context.Context

	snsActor *actions.SnsActor
	sqsActor *actions.SqsActor

	baseEndpoint string
	snsTopicArn  string
	sqsQueueUrl  string
	maxMessages  int
	waitTime     int
)

func init() {
	app = cli.NewApp()
	app.Name = "go sns example"
	app.Version = "0.0.0"

	baseEndpoint = "http://localhost:4566"
	snsTopicArn = "arn:aws:sns:us-east-1:000000000000:dummy-topic"
	sqsQueueUrl = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/dummy-queue"
	maxMessages = 10
	waitTime = 1

	ctx = context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	sqsActor = &actions.SqsActor{
		SqsClient: sqs.NewFromConfig(sdkConfig, func(o *sqs.Options) {
			o.BaseEndpoint = aws.String(baseEndpoint)
		}),
		QueueUrl:    sqsQueueUrl,
		MaxMessages: int32(maxMessages),
		WaitTime:    int32(waitTime),
	}
	snsActor = &actions.SnsActor{
		SnsClient: sns.NewFromConfig(sdkConfig, func(o *sns.Options) {
			o.BaseEndpoint = aws.String(baseEndpoint)
		}),
		TopicArn: snsTopicArn,
	}
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:  "push",
			Usage: "publishes a message to a sns topic",
			Action: func(c *cli.Context) error {
				publishMessage()
				return nil
			},
		},
		{
			Name:  "poll",
			Usage: "polls a message from the actions queue",
			Action: func(c *cli.Context) error {
				pollMessages()
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func publishMessage() {
	count := rand.Intn(maxMessages-1) + 1
	events, err := event.GenerateEvents(count)
	if err != nil {
		panic(err)
	}

	for _, e := range events {
		log.Printf("Event - %+v", e)
		messageBytes, err := proto.Marshal(e)
		if err != nil {
			panic(err)
		}
		encodedMessage := base64.StdEncoding.EncodeToString(messageBytes)
		err = snsActor.Publish(ctx, encodedMessage)
		if err != nil {
			panic(err)
		}
	}
}

func pollMessages() {
	sqsActor.Poll(ctx)
}
