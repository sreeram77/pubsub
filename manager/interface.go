package manager

import "github.com/sreeram77/pubsub/event"

type Manager interface {
	RegisterSubscriber() error
	Broadcast(event.Event) error
}
