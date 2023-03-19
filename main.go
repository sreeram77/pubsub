package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/publisher"
	"google.golang.org/grpc"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Debug("Listening for Registration...")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	publisher.RegisterPublisherServer(grpcServer, publisher.NewServer())

	grpcServer.Serve(lis)
}
