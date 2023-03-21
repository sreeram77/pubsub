package manager

import (
	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
)

type eventManager struct {
	connections map[string]([]chan event.Event)
}

func New() Manager {
	connections := make(map[string]([]chan event.Event))
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
	log.Info("Event:", e)

	return nil
}
