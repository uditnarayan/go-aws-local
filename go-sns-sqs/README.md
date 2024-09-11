# go-sns-sqs

This project aims at creating a pub sub design using aws sns and sqs. The producer publishes an event to sns topic. The 
consumer subscribes to that topic using sqs and listen to those events.

## how to run

Clean the project, removes any existing docker infrastructure and deletes the binary.
```shell
$ make clean
```

Initialize the project, downloads the go dependencies and creates the binary
```shell
$ make
```

Create the resources and initialize the consumer. This will create an sns topic, a sqs queue which subscribes 
to that sns topic. Then the consumer process will be initiated, which will listen to sqs queue and prints the messages
received.
```shell
$ make poll
```

Open a new terminal and push the events to sns. The producer will generate some random events and publish them to sns.
```shell
$ make push
```

You can now switch back to the previous terminal and see those messages printed by consumer.
