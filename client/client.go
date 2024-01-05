package main

import (
	"log"

	"github.com/nimeshjohari02/grpc-golang/train"
	"golang.org/x/net/context"
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
	user := &train.User{
		FirstName: "Nimesh",
		LastName: "Johari",
		EmailAddress: "nimeshjohari95@gmail.com",

	}
	removeUserResponse, err := client.RemoveUser(context.Background(), user)
	if err != nil {
		log.Fatalf("Unable to remove user %v", err)
	}
	log.Printf("User removed: %t", removeUserResponse.Success)
}
// Run the client
// go run client.go
