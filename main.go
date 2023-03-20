package main

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/sreeram77/pubsub/manager"
	"github.com/sreeram77/pubsub/publisher"
	"github.com/sreeram77/pubsub/subscriber"
	"google.golang.org/grpc"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.TraceLevel)

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Debug("Listening for Registration...")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	eventManager := manager.New()

	publisherServer := publisher.NewServer(eventManager)
	subscriberServer := subscriber.NewServer(eventManager)

	publisher.RegisterPublisherServer(grpcServer, publisherServer)
	subscriber.RegisterSubscriberServer(grpcServer, subscriberServer)

	grpcServer.Serve(lis)
}
