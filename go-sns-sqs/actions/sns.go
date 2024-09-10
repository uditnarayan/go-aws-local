package actions

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

type SnsActor struct {
	SnsClient *sns.Client
}

func (actor SnsActor) Publish(ctx context.Context, topicArn string, message string) error {
	publishInput := sns.PublishInput{TopicArn: aws.String(topicArn), Message: aws.String(message)}
	_, err := actor.SnsClient.Publish(ctx, &publishInput)
	if err != nil {
		log.Printf("Couldn't publish message to topic %v. Here's why: %v", topicArn, err)
		return err
	}
	return nil
}
