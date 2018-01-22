package main

import (
	"fmt"
	"sync"
	"time"

	mygrpc "github.com/sky0621/try-grpc-go/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var loopCount = 10000

// Counter ...
type Counter struct {
	mux *sync.RWMutex
	idx int
}

// CountUp ...
func (c *Counter) CountUp() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.idx = c.idx + 1
}

var idx = 0

func main() {
	fmt.Println("[main] Start")

	c := &Counter{
		mux: &sync.RWMutex{},
		idx: 0,
	}

	// part1(c)
	part2(c)

	fmt.Println("[main] End")
	fmt.Println(c.idx)
}

func part1(c *Counter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		go tryOnetime(c, wg)
	}
	wg.Wait()
}

func part2(c *Counter) {
	conn := createConn()
	if conn == nil {
		return
	}
	defer func() {
		time.Sleep(3 * time.Second)
		conn.Close()
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		go execCreateMessage(c, conn, wg)
	}
	wg.Wait()
}

func tryOnetime(c *Counter, wg *sync.WaitGroup) {
	conn := createConn()
	if conn == nil {
		return
	}
	defer func() {
		time.Sleep(3 * time.Second)
		conn.Close()
		wg.Done()
	}()
	client := mygrpc.NewDirectMessagesServiceClient(conn)
	_, err := client.CreateMessage(context.Background(), &mygrpc.CreateMessageRequest{
		MessageCreate: &mygrpc.MessageCreate{
			Target:      &mygrpc.Target{RecipientId: "recp123456"},
			MessageData: &mygrpc.MessageData{Text: "Hello, GRPC!"},
		},
	})
	if err == nil {
		// fmt.Println(res)
	} else {
		fmt.Printf("could not CreateMessage: %v\n", err)
	}

	c.CountUp()
}

func createConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err == nil {
		return conn
	}
	fmt.Printf("could not grpc.Dial: %v", err)
	return nil
}

func execCreateMessage(c *Counter, conn *grpc.ClientConn, wg *sync.WaitGroup) {
	defer wg.Done()
	client := mygrpc.NewDirectMessagesServiceClient(conn)
	_, err := client.CreateMessage(context.Background(), &mygrpc.CreateMessageRequest{
		MessageCreate: &mygrpc.MessageCreate{
			Target:      &mygrpc.Target{RecipientId: "recp123456"},
			MessageData: &mygrpc.MessageData{Text: "Hello, GRPC!"},
		},
	})
	if err == nil {
		// fmt.Println(res)
	} else {
		fmt.Printf("could not CreateMessage: %v\n", err)
	}

	c.CountUp()
}
