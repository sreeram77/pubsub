package subscriber

import (
	"context"

	"github.com/sreeram77/pubsub/manager"
)

type Server struct {
	manager *manager.Manager
}

func NewServer(m *manager.Manager) *Server {
	return &Server{manager: m}
}

func (s *Server) Subscribe(c context.Context, e *Topic) (*Ack, error) {
	return &Ack{
		Success: true,
		Message: "",
	}, nil
}

func (s *Server) mustEmbedUnimplementedSubscriberServer() {}
