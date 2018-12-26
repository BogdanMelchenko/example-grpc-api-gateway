package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/BogdanMelchenko/example/proto/logic"
	microservice "github.com/BogdanMelchenko/example/proto/microservice"
	"google.golang.org/grpc"
)

type ExampleServer struct {
	MicroserviceClient microservice.MicroServiceClient
}

func (s *ExampleServer) Echo(ctx context.Context, req *pb.StringMessage) (*pb.StringMessage, error) {
	res, _ := s.MicroserviceClient.GetString(context.Background(), &microservice.MicroMessage{})

	return &pb.StringMessage{Value: "hello there" + res.Value}, nil
}

func main() {
	conn, err := grpc.Dial(os.Getenv("MICROSERVICE_ADDRESS"), grpc.WithInsecure())
	//conn, err := grpc.Dial("microservice:9091", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	microclient := microservice.NewMicroServiceClient(conn)

	example := ExampleServer{microclient}

	flag.Parse()
	lis, err := net.Listen("tcp", os.Getenv("LOGIC_ADDRESS"))
	//lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterYourServiceServer(grpcServer, &example)

	grpcServer.Serve(lis)
}
