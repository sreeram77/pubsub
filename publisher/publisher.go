package publisher

import "context"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Publish(c context.Context, e *Event) (*Ack, error) {
	return &Ack{
		Success: true,
		Message: "",
	}, nil
}

func (s *Server) mustEmbedUnimplementedPublisherServer() {}
