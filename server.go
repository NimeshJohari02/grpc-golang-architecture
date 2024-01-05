package main

import (
	"log"
	"net"

	"github.com/nimeshjohari02/grpc-golang/train"
	"google.golang.org/grpc"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
	log.Fatalf("Unable to setup server %v" , err)
	} 
	s := train.Server{}
	grpcServer := grpc.NewServer()
	train.RegisterTrainTicketServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
	log.Fatalf("Failed to serve grpc Server at Port 9000 %v", err)
	}
}