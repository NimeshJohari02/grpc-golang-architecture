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
