init:
	go install
	protoc -I=. --go_out=. ./proto/event.proto

clean:
	docker-compose down
	rm -f go-sns-sqs
	rm -rf volume

aws:
	docker-compose up

poll:
	go run main.go poll

push:
	go run main.go push
