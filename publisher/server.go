package publisher

import (
	"io"

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

func (s *Server) mustEmbedUnimplementedPublisherServer() {}
