# go-sns-sqs

This project aims at creating a pub sub design using aws sns and sqs. The producer publishes an event to sns topic. The 
consumer subscribes to that topic using sqs and listen to those events. The events will be encoded using protobuf.

## how to run

Clean the project, removes any existing docker infrastructure and deletes the binary.
```shell
$ make clean
```

Initialize the project, downloads the go dependencies, build proto file and creates the binary
```shell
$ make
```

Initializes the localstack aws container. This will create a sns topic, a sqs queue which subscribes to that sns topic.
```shell
$ make aws
```

Open a new terminal and run the following command. It initializes the consumer which will listen to sqs queue and displays the messages published.
```shell
$ make poll
```

Open a new terminal and run the following command. The producer will generate some random events and publish them to sns.
```shell
$ make push
```

You can now switch back to the previous terminal and see those messages printed by consumer.
