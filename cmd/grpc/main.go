package main

import (
	"context"
	"net"

	"github.com/cpustejovsky/microservice"
	"github.com/cpustejovsky/microservice/internal/logger"
	pb "github.com/cpustejovsky/microservice/whitelist"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
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

func newServer() *WhiteListServer {
	s := &WhiteListServer{}
	return s
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWhiteListServer(s, &WhiteListServer{})
	if err := s.Serve(lis); err != nil {
		logger.Error.Fatalf("failed to serve: %v", err)
	}
}
