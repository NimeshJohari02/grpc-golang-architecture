package main

import (
	"log"

	"github.com/nimeshjohari02/grpc-golang/train"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
struct TicketsArray {
	tickets []*train.Ticket
}

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
	userTicket := &train.Ticket{
		User: user,
		From : "Bangalore",
		To: "Mumbai",
		PricePaid: 1000,
		Section: train.Section_SECTION_A,
		Seat : "13A",
	}
	confirmedTicket, err := client.SubmitPurchase(context.Background(), userTicket)
	if err != nil {
		log.Fatalf("Unable to purchase ticket %v", err)
	}
	log.Printf("Ticket confirmed for user %s %s. Ticket", confirmedTicket.User.FirstName, confirmedTicket.User.LastName)

	var (
		// Users
		alice = &train.User{
			FirstName:    "Alice",
			LastName:     "Anderson",
			EmailAddress: "",
		}
		bob = &train.User{
			FirstName:    "Bob",
			LastName:     "Baker",
			EmailAddress: "",
		}
		chris = &train.User{
			FirstName:    "Chris",
			LastName:     "Clark",
			EmailAddress: "",
		}
		// Tickets
		aliceTicket = &train.Ticket{
			User:      alice,
			From:      "Bangalore",
			To:        "Mumbai",
			PricePaid: 1000,
			Section:   train.Section_SECTION_A,
			Seat:      "13A",
		}
		bobTicket = &train.Ticket{
			User:      bob,
			From:      "Bangalore",
			To:        "Mumbai",
			PricePaid: 1000,
			Section:   train.Section_SECTION_A,
			Seat:      "13B",
		}
		chrisTicket = &train.Ticket{
			User:      chris,
			From:      "Bangalore",
			To:        "Mumbai",
			PricePaid: 1000,
			Section:   train.Section_SECTION_A,
			Seat:      "13C",
		}
		
	)
	// Submitting tickets
	aliceTicket, err = client.SubmitPurchase(context.Background(), aliceTicket)
	if err != nil {
		log.Fatalf("Unable to purchase ticket %v", err)
	}
	bobTicket, err = client.SubmitPurchase(context.Background(), bobTicket)
	if err != nil {
		log.Fatalf("Unable to purchase ticket %v", err)
	}
	chrisTicket, err = client.SubmitPurchase(context.Background(), chrisTicket)
	if err != nil {
		log.Fatalf("Unable to purchase ticket %v", err)
	}
	Traintickets := &TicketsArray{
		tickets: []*train.Ticket{
			aliceTicket,
			bobTicket,
			chrisTicket,
		},
	}
	removeUserResponse, err := client.RemoveUser(context.Background(), user)
	if err != nil {
		log.Fatalf("Unable to remove user %v", err)
	}
	log.Printf("User removed: %t", removeUserResponse.Success)
}
// Run the client
// go run client.go
