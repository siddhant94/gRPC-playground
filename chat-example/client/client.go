package client

import(
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/siddhant94/chat-grpc-example/chat"
)

func PingSayHelloServer(servAddr string)  {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(servAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed, did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello from client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}