package actions

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"go-sns-sqs/actions/event"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

type MessageBody struct {
	Message string
}

type SqsActor struct {
	SqsClient   *sqs.Client
	QueueUrl    string
	MaxMessages int32
	WaitTime    int32
}

func (actor SqsActor) GetMessages(ctx context.Context) []types.Message {
	result, err := actor.SqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(actor.QueueUrl),
		MaxNumberOfMessages: actor.MaxMessages,
		WaitTimeSeconds:     actor.WaitTime,
	})
	if err != nil {
		log.Printf("Failed to get messages from queue %v. Here's why: %v\n", actor.QueueUrl, err)
		return nil
	} else {
		return result.Messages
	}
}

func (actor SqsActor) DeleteMessages(ctx context.Context, messages []types.Message) {
	entries := make([]types.DeleteMessageBatchRequestEntry, len(messages))
	for msgIndex := range messages {
		entries[msgIndex].Id = aws.String(fmt.Sprintf("%v", msgIndex))
		entries[msgIndex].ReceiptHandle = messages[msgIndex].ReceiptHandle
	}
	_, err := actor.SqsClient.DeleteMessageBatch(ctx, &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(actor.QueueUrl),
	})
	if err != nil {
		log.Printf("Failed to delete messages from queue %v. Here's why: %v\n", actor.QueueUrl, err)
	}
}

func (actor SqsActor) Poll(ctx context.Context) {
	log.Printf("Starting polling on queue: %v\n", actor.QueueUrl)
	for {
		messages := actor.GetMessages(ctx)
		if messages == nil {
			time.Sleep(1000)
			continue
		}

		log.Printf("\n")
		log.Printf("Reading %d messages", len(messages))
		for _, message := range messages {
			e, err := parseAndDecodeMessage(message)
			if err != nil {
				log.Printf("Failed to process message. Here's why: %v\n", err)
			}
			log.Printf("Message %v: %v\n", *message.MessageId, e.ToString())
		}
		actor.DeleteMessages(ctx, messages)
		log.Printf("Deleted %d messages", len(messages))
	}
}

func parseAndDecodeMessage(message types.Message) (event.Event, error) {
	messageBody := MessageBody{}
	var e event.Event

	err := json.Unmarshal([]byte(*message.Body), &messageBody)
	if err != nil {
		log.Printf("Failed to parse json message.")
		return e, err
	}
	decodedMsg, err := base64.StdEncoding.DecodeString(messageBody.Message)
	if err != nil {
		log.Fatalf("Failed to decode base64 message.")
		return e, err
	}

	err = proto.Unmarshal(decodedMsg, &e)
	if err != nil {
		log.Printf("Failed to unmarshall protobuf message.")
		return e, err
	}
	return e, nil
}
