package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/micro-company/go-mail-crd/grpc/mail"
	"net"
	"log"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SendMail(ctx context.Context, in *pb.MailRequest) (*pb.MailResponse, error) {
	return &pb.MailResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

    // Create a new gRPC server
    s := grpc.NewServer()
    pb.RegisterMailServer(s, &server{})
    if err := s.Serve(lis); err != nil {
    	log.Fatalf("failed to server: %v", err)
	}
}