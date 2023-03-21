package publisher

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/event"
	"github.com/sreeram77/pubsub/manager"
)

// Server is responsible for serving publisher requests.
type Server struct {
	manager manager.Manager
}

// NewServer returns a new instance of publisher Server.
func NewServer(m manager.Manager) *Server {
	return &Server{manager: m}
}

// Publish polls for published events and forwards it to event manager for
// broadcasting to the subscribers of that topic.
func (s *Server) Publish(stream Publisher_PublishServer) error {
	for {
		e, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&Status{
					Success: true,
					Message: "",
				})
			}

			log.Error(err)
			return err
		}

		s.manager.Broadcast(event.Event{
			Topic: e.GetTopic(),
			Body:  e.GetMessage(),
		})
	}
}

// mustEmbedUnimplementedPublisherServer is added to implement PublisherServer interface.
func (s *Server) mustEmbedUnimplementedPublisherServer() {}
