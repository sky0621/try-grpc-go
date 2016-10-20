package main

import (
	"log"
	"time"

	trygrpcgo "github.com/sky0621/try-grpc-go"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		time.Sleep(10 * time.Second)
		conn.Close()
	}()

	client := trygrpcgo.NewPersonSenderClient(conn)
	p := trygrpcgo.Person{
		Name:   "Taro",
		Id:     1234567890,
		Email:  "hoge@example.com",
		Phones: createSamplePhones(),
	}
	r, err := client.SendPerson(context.Background(), &p)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}

func createSamplePhones() []*trygrpcgo.Person_PhoneNumber {
	var samplePhone1 = trygrpcgo.Person_PhoneNumber{Number: "080-5948-7404", Type: trygrpcgo.Person_MOBILE}
	var samplePhone2 = trygrpcgo.Person_PhoneNumber{Number: "03-1234-5678", Type: trygrpcgo.Person_HOME}
	var samplePhones []*trygrpcgo.Person_PhoneNumber
	samplePhones = append(samplePhones, &samplePhone1)
	samplePhones = append(samplePhones, &samplePhone2)
	return samplePhones
}
