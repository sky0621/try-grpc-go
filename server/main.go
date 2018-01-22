package main

import (
	"fmt"
	"log"
	"net"

	context "golang.org/x/net/context"

	mygrpc "github.com/sky0621/try-grpc-go/grpc"

	"google.golang.org/grpc"
)

// DirectMessagesServiceServerImpl ...
type DirectMessagesServiceServerImpl struct{}

// CreateMessage ...
func (s DirectMessagesServiceServerImpl) CreateMessage(ctx context.Context, req *mygrpc.CreateMessageRequest) (*mygrpc.CreateMessageResponse, error) {
	fmt.Printf("Target: %#v\n", req.MessageCreate.Target)
	fmt.Printf("MessageData: %#v\n", req.MessageCreate.MessageData)
	return &mygrpc.CreateMessageResponse{
		Event: &mygrpc.Event{
			Id:               "1234567890",
			CreatedTimestamp: "20180122",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5050))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mygrpc.RegisterDirectMessagesServiceServer(grpcServer, &DirectMessagesServiceServerImpl{})
	grpcServer.Serve(lis)
}
