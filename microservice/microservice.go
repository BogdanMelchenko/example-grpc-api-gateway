package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	pb "github.com/BogdanMelchenko/example/proto/microservice"
	"google.golang.org/grpc"
)

type MicroservicServer struct {
}

func (s *MicroservicServer) GetString(ctx context.Context, req *pb.MicroMessage) (*pb.MicroMessage, error) {
	return &pb.MicroMessage{Value: "Generale Kenobi"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", os.Getenv("MICROSERVICE_ADDRESS"))
	//lis, err := net.Listen("tcp", ":9091")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMicroServiceServer(grpcServer, &MicroservicServer{})

	grpcServer.Serve(lis)
}
