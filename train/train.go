package train

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

// ModifySeat implements TrainTicketServiceServer.
func (*Server) ModifySeat(context.Context, *ModifySeatRequest) (*Ticket, error) {
	panic("unimplemented")
}

// SubmitPurchase implements TrainTicketServiceServer.
func (*Server) SubmitPurchase(context.Context, *Ticket) (*Ticket, error) {
	panic("unimplemented")
}

// ViewReceipt implements TrainTicketServiceServer.
func (*Server) ViewReceipt(context.Context, *User) (*Ticket, error) {
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

func (s *Server) RemoveUser(ctx context.Context, user *User) (*RemoveUserResponse, error) {
	log.Println("RemoveUser called with user: ", user.FirstName, user.LastName)
	return &RemoveUserResponse{Success: true}, nil
}
