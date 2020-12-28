package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/cpustejovsky/microservice/whitelist"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWhiteListClient(conn)
	fmt.Println(c)
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
		log.Fatalf("Could not verify if IP address is whitelisted: %v", err)
	}
	if r.GetWhiteListed() {
		log.Printf("IP Address %v is whitelisted", in.IP)
	}
}
