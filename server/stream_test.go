package server

import (
	"google.golang.org/grpc"
	pb "grpc/proto/stream"
	"log"
	"net"
	"testing"
)

func TestStreamServer(t *testing.T) {
	server := grpc.NewServer() //创建 gRPC Server 对象
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
