# pubsub

pubsub is a light-weight publisher-subscriber service using gRPC as the medium of communication. This acts as an event bus that delivers events sent by the producers to the respective subscribers who have registered for a topic.

## Producer
The producer can push events with a topic and a message. Publisher should use the proto in /publisher/publisher.proto for implementation.

## Subscriber
The subscriber can subsribe to a topic and listen to events for that topic. Subscriber should use the proto in /subscriber/subscriber.proto for implementation.

## Installation 
```sh
go run main.go
```

## Roadmap
- Replace map with in-memory database to support horizontal scaling.
- Persist subscription information.
