init:
	go build

clean:
	docker-compose down
	rm -f go-sns-sqs

poll:
	docker-compose up -d
	sh -f init.sh
	./go-sns-sqs poll

push:
	./go-sns-sqs push
