package storage

import "github.com/sreeram77/pubsub/event"

type Cache interface {
	Get(string) []chan event.Event
	Set(string, chan event.Event)
}
