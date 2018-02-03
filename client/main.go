package main

import (
	"fmt"
	"sync"
	"time"

	mygrpc "github.com/sky0621/try-grpc-go/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
)

var loopCount = 50000

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

	actionWithEachConnect(c)
	// actionWithCommonConnect(c)

	fmt.Println("[main] End")
	fmt.Println(c.idx)
}

// 各goroutineでの実行の都度、コネクションを生成
func actionWithEachConnect(c *Counter) {
	wg := &sync.WaitGroup{}
	for i := 0; i < loopCount; i++ {
		wg.Add(1)
		go func(c *Counter, wg *sync.WaitGroup) {
			conn := createConn() // ★goroutine内でコネクション生成
			if conn == nil {
				return
			}
			defer func() {
				time.Sleep(3 * time.Second)
				conn.Close()
				wg.Done()
			}()
			err := createMessage(conn)
			if err != nil {
				fmt.Printf("could not CreateMessage: %v\n", err)
			}

			c.CountUp()
		}(c, wg)
	}
	wg.Wait()
}

// 事前に生成済みのコネクションを各goroutineに渡して実行
func actionWithCommonConnect(c *Counter) {
	conn := createConn() // ★goroutine実行前にコネクションを生成
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
		go func(c *Counter, conn *grpc.ClientConn, wg *sync.WaitGroup) {
			defer wg.Done()
			err := createMessage(conn)
			if err != nil {
				fmt.Printf("could not CreateMessage: %v\n", err)
			}

			c.CountUp()
		}(c, conn, wg)
	}
	wg.Wait()
}

func createConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err == nil {
		return conn
	}
	fmt.Printf("could not grpc.Dial: %v", err)
	return nil
}

func createMessage(conn *grpc.ClientConn) error {
	client := mygrpc.NewDirectMessagesServiceClient(conn)
	_, err := client.CreateMessage(context.Background(), &mygrpc.CreateMessageRequest{
		MessageCreate: &mygrpc.MessageCreate{
			Target:      &mygrpc.Target{RecipientId: "recp123456"},
			MessageData: &mygrpc.MessageData{Text: "Hello, GRPC!"},
		},
	})
	return err
}
