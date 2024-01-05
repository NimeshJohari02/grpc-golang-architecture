package main

import (
	"log"

	"github.com/nimeshjohari02/grpc-golang/train"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn , err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Unable to connect to server %v", err)
	}
	defer conn.Close()
	client := train.NewTrainTicketServiceClient(conn)
	if err != nil {
		log.Fatalf("Unable to purchase ticket %v", err)
	}
}
// Run the client
// go run client.go
