package main

import (
	"context"
	"time"

	"github.com/cpustejovsky/microservice/internal/logger"

	pb "github.com/cpustejovsky/microservice/whitelist"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Error.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWhiteListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	in := pb.Input{
		IP: "81.2.69.142",
		WhiteList: []string{
			"United Kingdom",
			"United States",
			"Mexico",
		},
	}
	r, err := c.CheckIPAddress(ctx, &in)
	if err != nil {
		logger.Error.Fatalf("Could not verify if IP address is whitelisted: %v", err)
	}
	if r.GetWhiteListed() {
		logger.Info.Printf("IP Address %v is whitelisted", in.IP)
	}
}
