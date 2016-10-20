package main

import (
	"log"

	trygrpcgo "github.com/sky0621/try-grpc-go"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := trygrpcgo.NewPersonSenderClient(conn)
	p := trygrpcgo.Person{Name: "Who?"}
	r, err := client.SendPerson(context.Background(), &p)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
