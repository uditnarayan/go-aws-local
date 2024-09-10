package actions

import (
	"github.com/go-faker/faker/v4"
	"log"
)

type Event struct {
	Id        string `faker:"uuid_hyphenated"`
	Timestamp string `faker:"timestamp"`
	Username  string `faker:"username"`
	FirstName string `faker:"first_name"`
	LastName  string `faker:"last_name"`
}

func GenerateEvents(count int) ([]*Event, error) {
	events := make([]*Event, count)
	for i := 0; i < count; i++ {
		event := &Event{}
		err := faker.FakeData(event)
		if err != nil {
			return nil, err
		}
		events[i] = event
	}
	log.Printf("Generated %d events", count)
	return events, nil
}
