package main

import (
	"log"
	"time"

	mygrpc "github.com/sky0621/try-grpc-go/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		time.Sleep(3 * time.Second)
		conn.Close()
	}()

	client := mygrpc.NewDirectMessagesServiceClient(conn)
	res, err := client.CreateMessage(context.Background(), &mygrpc.CreateMessageRequest{
		MessageCreate: &mygrpc.MessageCreate{
			Target:      &mygrpc.Target{RecipientId: "recp123456"},
			MessageData: &mygrpc.MessageData{Text: "Hello, GRPC!"},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Event)
}
