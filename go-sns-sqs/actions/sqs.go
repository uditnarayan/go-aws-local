package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"log"
)

type SqsActor struct {
	SqsClient *sqs.Client
}

type MessageBody struct {
	Message string
}

func (actor SqsActor) GetMessages(ctx context.Context, queueUrl string, maxMessages int32, waitTime int32) []types.Message {
	result, err := actor.SqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: maxMessages,
		WaitTimeSeconds:     waitTime,
	})
	if err != nil {
		log.Printf("Couldn't get messages from queue %v. Here's why: %v\n", queueUrl, err)
		return nil
	} else {
		return result.Messages
	}
}

func (actor SqsActor) DeleteMessages(ctx context.Context, queueUrl string, messages []types.Message) {
	entries := make([]types.DeleteMessageBatchRequestEntry, len(messages))
	for msgIndex := range messages {
		entries[msgIndex].Id = aws.String(fmt.Sprintf("%v", msgIndex))
		entries[msgIndex].ReceiptHandle = messages[msgIndex].ReceiptHandle
	}
	_, err := actor.SqsClient.DeleteMessageBatch(ctx, &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(queueUrl),
	})
	if err != nil {
		log.Printf("Couldn't delete messages from queue %v. Here's why: %v\n", queueUrl, err)
	}
}

func (actor SqsActor) Poll(ctx context.Context, queueUrl string, maxMessages int32, waitTime int32) {
	log.Printf("Starting polling on queue: %v\n", queueUrl)
	for {
		messages := actor.GetMessages(ctx, queueUrl, maxMessages, waitTime)
		if messages == nil {
			continue
		}

		log.Printf("\n")
		log.Printf("Reading %d messages", len(messages))
		for _, message := range messages {
			messageBody := MessageBody{}
			err := json.Unmarshal([]byte(*message.Body), &messageBody)
			if err != nil {
				log.Printf("Couldn't parse json message. Here's why: %v\n", err)
			}
			log.Printf("Message %v: %v\n", message.MessageId, messageBody.Message)
		}
		actor.DeleteMessages(ctx, queueUrl, messages)
		log.Printf("Deleted %d messages", len(messages))
	}
}
