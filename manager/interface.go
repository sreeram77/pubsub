package manager

import (
	"github.com/sreeram77/pubsub/event"
)

type Manager interface {
	RegisterSubscriber(string, chan event.Event) error
	Broadcast(event.Event) error
}
