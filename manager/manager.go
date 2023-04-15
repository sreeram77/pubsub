package manager

import (
	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
	"github.com/sreeram77/pubsub/storage"
)

type subscribers []chan event.Event

type eventManager struct {
	connections storage.Cache
}

// New returns an instance of eventManager that implements Manager interface.
func New(r storage.Cache) Manager {
	return &eventManager{
		connections: r,
	}
}

func (em *eventManager) RegisterSubscriber(topics []string, eventChan chan event.Event) error {
	log.Info("Registering subscriber for topic:", topics)

	// Save topics with subscriber connection
	for _, topic := range topics {
		em.connections.Set(topic, eventChan)
	}

	return nil
}

func (em *eventManager) Broadcast(e event.Event) error {
	log.Info("Broadcasting event:", e)

	// Get all connections for the topic
	conns := em.connections.Get(e.Topic)

	// Broadcast event to all subscribers
	for _, conn := range conns {
		go func(c chan event.Event) {
			c <- e
		}(conn)
	}

	return nil
}
