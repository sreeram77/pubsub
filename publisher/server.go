package publisher

import (
	"context"

	"github.com/sreeram77/pubsub/manager"
)

type Server struct {
	manager manager.Manager
}

func NewServer(m manager.Manager) *Server {
	return &Server{manager: m}
}

func (s *Server) Publish(c context.Context, e *Event) (*Ack, error) {
	return &Ack{
		Success: true,
		Message: "",
	}, nil
}

func (s *Server) mustEmbedUnimplementedPublisherServer() {}
