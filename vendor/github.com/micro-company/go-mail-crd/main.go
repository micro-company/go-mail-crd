package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/micro-company/go-mail-crd/handlers/mail"
	"errors"


	pb "github.com/micro-company/go-mail-crd/grpc/mail"
	"net"
)

const (
	port = ":50051"
)

var (
	log = logrus.New()
)

type server struct{}

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}


func (s *server) SendMail(ctx context.Context, in *pb.MailRequest) (*pb.MailResponse, error) {
	// Send email
	data := mail.RecoveryData{
		Mail: in.Mail,
		Url:  in.Url,
	}
	err := mail.Recovery(data)
	if err != nil {
		return &pb.MailResponse{Success: false}, errors.New("failed to send message")
	}

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
	log.Info("Run services on port " + port)
    if err := s.Serve(lis); err != nil {
    	log.Fatalf("failed to server: %v", err)
	}
}