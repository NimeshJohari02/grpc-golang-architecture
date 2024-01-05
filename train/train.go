package train

// Import MongoDb driver libraries
import (
	"context"
	"log"
)
type Server struct {
}

var TicketsArray []*Ticket = make([]*Ticket, 0)
var globalUserArray []*User = make([]*User, 0)
var SectionAUsers []*User = make([]*User, 0)
var SectionBUsers []*User = make([]*User, 0)
// PurchaseTicket implements TrainTicketServiceServer.
func (*Server) PurchaseTicket(c context.Context, req *PurchaseTicketRequest) (*Ticket, error) {
	created_ticket := &Ticket{
		From : req.From,
		To : req.To,
		UserId : req.User.UserId,
		PriceOfTicket: req.PriceOfTicket,
		Section : req.Section,
		Seat : "9A",
		TrainId: req.TrainId,
	}
	if req.Section == 0 {
		SectionAUsers = append(SectionAUsers, req.User)
	} else {
		SectionBUsers = append(SectionBUsers, req.User)
	}

	globalUserArray = append(globalUserArray, req.User);
	log.Println(created_ticket , "created_ticket")
	TicketsArray = append(TicketsArray , created_ticket)
	return created_ticket , nil
}

// RemoveUserFromTrain implements TrainTicketServiceServer.
func (*Server) RemoveUserFromTrain(c context.Context, req *RemoveUserFromTrainRequest) (*RemoveUserResponse, error) {
	userId := req.UserId;
	trainId := req.TrainId;
	for i := 0; i < len(TicketsArray); i++ {
		if TicketsArray[i].UserId == userId && TicketsArray[i].TrainId == trainId {
			TicketsArray = append(TicketsArray[:i], TicketsArray[i+1:]...)
			return &RemoveUserResponse{Success: true}, nil
		}
	}
	return &RemoveUserResponse{Success: false}, nil

}

// ViewReceipt implements TrainTicketServiceServer.
func (*Server) ViewReceipt(ctx context.Context, req *ViewReceiptRequest) (*Ticket, error) {
	for i := 0; i < len(TicketsArray); i++ {
		if TicketsArray[i].UserId == req.UserId {
			return TicketsArray[i], nil
		}
	}
	return nil, nil
}

// ViewUsersBySection implements TrainTicketServiceServer.
func (*Server) ViewUsersBySection(ctx context.Context, req  *ViewUsersBySectionRequest) (*ViewUsersBySectionResponse, error) {
	section := req.Section;
	if section == 0 {
		return &ViewUsersBySectionResponse{Users: SectionAUsers}, nil
	} else {
		return &ViewUsersBySectionResponse{Users: SectionBUsers}, nil
	}
}
func (*Server) ModifySeat(ctx context.Context, req  *ModifySeatRequest) (*WasModifiedResponse, error) {
	UserID := req.UserId;
	TrainID := req.TrainId;
	NewSection := req.NewSection;
	NewSeat := req.NewSeat;
	for i := 0; i < len(TicketsArray); i++ {
		if TicketsArray[i].UserId == UserID && TicketsArray[i].TrainId == TrainID {
			TicketsArray[i].Section = NewSection;
			// Check for the change in section and update the global array
			TicketsArray[i].Seat = NewSeat;
			return &WasModifiedResponse{Success: true}, nil
		}
	}
}
// mustEmbedUnimplementedTrainTicketServiceServer implements TrainTicketServiceServer.
func (*Server) mustEmbedUnimplementedTrainTicketServiceServer() {
	panic("unimplemented")
}
