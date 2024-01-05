package train

// Import MongoDb driver libraries
import (
	"context"
)

type Server struct {
}


// ModifySeat implements TrainTicketServiceServer.
func (*Server) ModifySeat(context.Context, *ModifySeatRequest) (*WasModifiedResponse, error) {
	panic("unimplemented")
}

// PurchaseTicket implements TrainTicketServiceServer.
func (*Server) PurchaseTicket(context.Context, *PurchaseTicketRequest) (*Ticket, error) {
	// destucture the PurchaseTicketRequest to get User and Train
	// Create a new ticket
	// Save the ticket to the database
	// return the ticket
	// Destructure from field from PurchaseTicketRequest
	created_ticket := &Ticket{
		From : PurchaseTicketRequest.From,
		To : &PurchaseTicketRequest.To,
		UserId : &PurchaseTicketRequest.User.user_id,
		PriceOfTicket: &PurchaseTicketRequest.PriceOfTicket,
		Section : &PurchaseTicketRequest.Section,
		Seat : &PurchaseTicketRequest.Seat,
		TrainId: &PurchaseTicketRequest.train_id,
	}
	panic("unimplemented")
}

// RemoveUserFromTrain implements TrainTicketServiceServer.
func (*Server) RemoveUserFromTrain(context.Context, *RemoveUserFromTrainRequest) (*RemoveUserResponse, error) {
	panic("unimplemented")
}

// ViewReceipt implements TrainTicketServiceServer.
func (*Server) ViewReceipt(context.Context, *ViewReceiptRequest) (*Ticket, error) {
	panic("unimplemented")
}

// ViewUsersBySection implements TrainTicketServiceServer.
func (*Server) ViewUsersBySection(context.Context, *ViewUsersBySectionRequest) (*ViewUsersBySectionResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedTrainTicketServiceServer implements TrainTicketServiceServer.
func (*Server) mustEmbedUnimplementedTrainTicketServiceServer() {
	panic("unimplemented")
}
