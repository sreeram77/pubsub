package event

import (
	"testing"
)

func TestNew(t *testing.T) {
	topic := "Test"
	message := "test subject"
	body := []byte(message)

	e := New(topic, body)

	if string(e.body) != message || e.topic != topic {
		t.Fatal("Event creation failed")
	}
}
