package event

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"log"
	"time"
)

func GenerateEvents(count int) ([]*Event, error) {
	events := make([]*Event, count)
	for i := 0; i < count; i++ {
		event := &Event{}
		event.Id = gofakeit.UUID()
		event.EventName = "user_created"
		event.Timestamp = time.Now().UTC().String()
		event.Payload = &Payload{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
		}
		events[i] = event
	}
	log.Printf("Generated %d events", count)
	return events, nil
}

func (e Event) ToString() string {
	return fmt.Sprintf("{id: %v, eventName: %v, timestamp: %v, username: %v, firstName: %v, lastName: %v}",
		e.GetId(), e.GetEventName(), e.GetTimestamp(), e.GetPayload().GetUsername(), e.GetPayload().GetFirstName(), e.GetPayload().GetLastName())
}
