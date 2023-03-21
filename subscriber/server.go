package subscriber

import (
	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
	"github.com/sreeram77/pubsub/manager"
)

type Server struct {
	manager manager.Manager
}

func NewServer(m manager.Manager) *Server {
	return &Server{manager: m}
}

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

func (s *Server) mustEmbedUnimplementedSubscriberServer() {}
