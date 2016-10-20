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

func (s pSendServer) SendPerson(context.Context, *trygrpcgo.Person) (*trygrpcgo.SendPersonReply, error) {
	log.Println("[SendPerson]START")
	// log.Println(p.Name)
	return &trygrpcgo.SendPersonReply{Message: "send OK !?"}, nil
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
