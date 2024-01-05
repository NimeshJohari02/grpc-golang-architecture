package train

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	
}
func(s *Server) RemoveUser(ctx context.Context, user *User) (*RemoveUserResponse, error) {
	log.Println("RemoveUser called with user: ", user.FirstName, user.LastName)
	return &RemoveUserResponse{Success: true}, nil
}