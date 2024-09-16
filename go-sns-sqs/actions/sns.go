package actions

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

type SnsActor struct {
	SnsClient *sns.Client
	TopicArn  string
}

func (actor SnsActor) Publish(ctx context.Context, message string) error {
	publishInput := &sns.PublishInput{TopicArn: aws.String(actor.TopicArn), Message: aws.String(message)}
	_, err := actor.SnsClient.Publish(ctx, publishInput)
	if err != nil {
		log.Printf("Failed to publish message to topic %v. Here's why: %v", actor.TopicArn, err)
		return err
	}
	return nil
}
