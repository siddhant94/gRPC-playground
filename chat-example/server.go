package main

import(
	"fmt"
	"os"
	"log"
	"net"
	"sync"

	"github.com/siddhant94/chat-grpc-example/chat"
	"github.com/siddhant94/chat-grpc-example/client"
	"google.golang.org/grpc"
)

func main() {
	log.SetOutput(os.Stdout)
	fmt.Println("Setting up a grpcServer")
	var wg sync.WaitGroup
	wg.Add(1)
	port := ":15000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	fmt.Println("Registering our Chat service")
	// Get our chat service
	s := chat.Server{}
	// Register our service
	chat.RegisterChatServiceServer(grpcServer, &s)

	// Launch a go routine
	go func ()  {
		defer wg.Done()
		fmt.Println("In our goroutine")
		// spawn a grpc server to listen to grpc client calls.
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	fmt.Println("Waiting for server to be up")
	// Wait for server to start before sending client service call.
	// wg.Wait()

	fmt.Println("Initiating client Request to sayHello grpc service")
	client.PingSayHelloServer(port)

}