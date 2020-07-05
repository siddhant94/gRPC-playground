package chat

import (
	"log"
	"golang.org/x/net/context"
)

type Server struct {

}

func(s *Server) SayHello(ctx context.Context, inp *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", inp.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}