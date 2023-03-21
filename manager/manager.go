package manager

import (
	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
)

type eventManager struct {
}

func New() Manager {
	return &eventManager{}
}

func (em *eventManager) RegisterSubscriber(topic string, eventChan chan event.Event) error {
	return nil
}

func (em *eventManager) Broadcast(e event.Event) error {
	log.Info("Event:", e)

	return nil
}
