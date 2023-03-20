package event

import "time"

type Event struct {
	Topic      string
	Body       []byte
	receivedAt time.Time
}

// New returns a new instance of Event
func New(topic string, body []byte) Event {
	return Event{
		Topic: topic,
		Body:  body,
	}
}

// SetReceivedAt sets the receivedAt time for the event.
func (e *Event) SetReceivedAt(t time.Time) {
	e.receivedAt = t
}

func (e *Event) SetNowAsReceivedAt() {
	e.receivedAt = time.Now()
}
