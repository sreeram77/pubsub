package storage

import "github.com/sreeram77/pubsub/event"

type subscribers []chan event.Event

type cache struct {
	connections map[string]subscribers
}

func New() Cache {
	connections := make(map[string]subscribers)

	return &cache{
		connections: connections,
	}
}

func (c *cache) Get(key string) []chan event.Event {
	if value, found := c.connections[key]; found {
			return value
	}
	
	return nil
}

func (c *cache) Set(key string, value chan event.Event) {
	_, found := c.connections[key]
	if found {
		c.connections[key] = append(c.connections[key], value)
		return
	}

	c.connections[key] = []chan event.Event{value}
}
