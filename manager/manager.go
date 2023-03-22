package manager

import (
	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
)

type subscribers []chan event.Event

type eventManager struct {
	connections map[string]subscribers
}

// New returns an instance of eventManager that implements Manager interface.
func New() Manager {
	connections := make(map[string]subscribers)
	return &eventManager{
		connections: connections,
	}
}


func (em *eventManager) RegisterSubscriber(topic string, eventChan chan event.Event) error {
	log.Info("Registering subscriber for topic:", topic)

	_, found := em.connections[topic]
	if found {
		em.connections[topic] = append(em.connections[topic], eventChan)
		return nil
	}

	em.connections[topic] = []chan event.Event{eventChan}

	return nil
}

func (em *eventManager) Broadcast(e event.Event) error {
	log.Info("Broadcasting event:", e)

	if value, found := em.connections[e.Topic]; found {
		broadcastEvent(e, value)
	}

	return nil
}

func broadcastEvent(e event.Event, s []chan event.Event) {
	for _, v := range s {
		go func(c chan event.Event) {
			c <- e
		}(v)
	}
}
