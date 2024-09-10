# configure a test profile
aws configure set aws_access_key_id "dummy" --profile test-profile
aws configure set aws_secret_access_key "dummy" --profile test-profile
aws configure set region "us-east-1" --profile test-profile
aws configure set output "table" --profile test-profile

# Create an SNS topic
aws --endpoint-url=http://localhost:4566 sns create-topic --name dummy-topic --region us-east-1 --profile test-profile --output table | cat

# Create an SQS queue
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name dummy-queue --profile test-profile --region us-east-1 --output table | cat

# Add Subscription
aws --endpoint-url=http://localhost:4566 sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:dummy-topic --profile test-profile --region us-east-1 --protocol sqs --notification-endpoint arn:aws:sqs:us-east-1:000000000000:dummy-queue --output table | cat