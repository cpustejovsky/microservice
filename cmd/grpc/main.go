package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/cpustejovsky/microservice"
	pb "github.com/cpustejovsky/microservice/whitelist"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
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
		Error:       err.Error(),
	}, nil
}

func newServer() *WhiteListServer {
	s := &WhiteListServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWhiteListServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
