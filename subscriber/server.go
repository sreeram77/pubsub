package subscriber

import "context"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Subscribe(c context.Context, e *Topic) (*Ack, error) {
	return &Ack{
		Success: true,
		Message: "",
	}, nil
}

func (s *Server) mustEmbedUnimplementedSubscriberServer() {}
