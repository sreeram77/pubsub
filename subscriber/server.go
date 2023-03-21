package subscriber

import (
	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
	"github.com/sreeram77/pubsub/manager"
)

// Server is responsible for serving subscriber requests.
type Server struct {
	manager manager.Manager
}

// NewServer returns a new instance of subscriber Server.
func NewServer(m manager.Manager) *Server {
	return &Server{manager: m}
}

// Subscribe registers the subscriber to a topic.
// It responds to the subscriber when an event matching the topic is received.
func (s *Server) Subscribe(t *Topic, stream Subscriber_SubscribeServer) error {
	log.Info("subscribed to topic:", t.GetTopic())
	e := make(chan event.Event, 100)
	s.manager.RegisterSubscriber(t.GetTopic(), e)

	for {
		subEvent := <-e

		stream.Send(&REvent{
			Topic:   subEvent.Topic,
			Message: subEvent.Body,
		})
	}

}

// mustEmbedUnimplementedSubscriberServer is added to implement SubscriberServer interface.
func (s *Server) mustEmbedUnimplementedSubscriberServer() {}
