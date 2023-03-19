package event

import "time"

type Event struct {
	topic      string
	body       []byte
	receivedAt time.Time
}

// New returns a new instance of Event
func New(topic string, body []byte) Event {
	return Event{
		topic: topic,
		body: body,
	}
}

// Topic returns the topic of the event.
func (e *Event) Topic() string {
	return e.topic
}

// Body returns the body of the event.
func (e *Event) Body() []byte {
	return e.body
}

// SetReceivedAt sets the receivedAt time for the event.
func (e *Event) SetReceivedAt(t time.Time) {
	e.receivedAt = t
}

func (e *Event) SetNowAsReceivedAt() {
	e.receivedAt = time.Now()
}
