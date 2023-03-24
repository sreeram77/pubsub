# pubsub

pubsub is a light-weight publisher-subscriber service using gRPC as the medium of communication. This acts as an event bus that delivers events sent by the producers to the respective subscribers who have registered for a topic.

## Installation 
```sh
go run main.go
```

## Roadmap
- Support subscription of multiple events.
- Replace map with in-memory database to support horizontal scaling.
- Persist subscription information.
