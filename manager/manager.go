package manager

import "github.com/sreeram77/pubsub/event"

type eventManager struct {
}

func New() Manager {
	return &eventManager{}
}

func (e *eventManager) RegisterSubscriber() error {
	return nil
}

func (e *eventManager) Broadcast(event.Event) error {
	return nil
}
