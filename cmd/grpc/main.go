package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/pkg/logger"
	pb "github.com/cpustejovsky/microservice/whitelist"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type WhiteListServer struct {
	pb.UnimplementedWhiteListServer
}

func (s *WhiteListServer) CheckIPAddress(ctx context.Context, Input *pb.Input) (*pb.Output, error) {
	ok, err := microservice.CheckIPAddress(Input.IP, Input.WhiteList)
	if err != nil || !ok {
		return nil, err
	}
	return &pb.Output{
		WhiteListed: ok,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		logger.Error.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWhiteListServer(s, &WhiteListServer{})
	if err := s.Serve(lis); err != nil {
		logger.Error.Fatalf("failed to serve: %v", err)
	}
}
