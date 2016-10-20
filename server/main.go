package main

import (
	"flag"
	"log"
	"net"

	context "golang.org/x/net/context"

	trygrpcgo "github.com/sky0621/try-grpc-go"

	"google.golang.org/grpc"
)

type pSendServer struct{}

func (s pSendServer) SendPerson(ctx context.Context, p *trygrpcgo.Person) (*trygrpcgo.SendPersonReply, error) {
	log.Println("[SendPerson]START")
	log.Println(p.Name)
	log.Println(p.Id)
	log.Println(p.Email)
	for _, phone := range p.Phones {
		log.Println(phone)
	}
	return &trygrpcgo.SendPersonReply{Message: "OK! We[server] are received your data."}, nil
}

var server trygrpcgo.PersonSenderServer

func main() {
	addrFlag := flag.String("p", ":5050", "host:port")
	lis, err := net.Listen("tcp", *addrFlag)
	if err != nil {
		log.Fatal(err)
	}

	var svr trygrpcgo.PersonSenderServer
	var pSvr pSendServer
	svr = pSvr

	grpcServer := grpc.NewServer()
	trygrpcgo.RegisterPersonSenderServer(grpcServer, svr)
	grpcServer.Serve(lis)
}
