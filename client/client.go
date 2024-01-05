package main

import (
	"context"
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
	confirmed_ticket , err := client.PurchaseTicket(context.Background(), &train.PurchaseTicketRequest{
		From: "New York",
		To: "Boston",
		User: &train.User{
			UserId: "4",
			FirstName: "Gaurav",
			EmailAddress: "guarav@gmail.com",
			HasTicket: false,
		},
		PriceOfTicket: 10043,
		Section: train.Section_SECTION_B,
		TrainId: "1",
	})
	if err != nil {
		log.Fatalf("Unable to purchase ticket %v", err)
	}
	log.Println(confirmed_ticket , "confirmed_ticket")
	
	// Modify Seat
	modified_seat , err := client.ModifySeat(context.Background(), &train.ModifySeatRequest{
		UserId: "4",
		TrainId: "1",
		Section: train.Section_SECTION_B,
		NewSection: train.Section_SECTION_A,
		NewSeat: "9A",

	})
	if err != nil {
		log.Fatalf("Unable to modify seat %v", err)
	}
	log.Println(modified_seat , "modified_seat")
	
}
// Run the client
// go run client.go
