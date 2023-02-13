package server

import (
	"google.golang.org/grpc"
	pb "grpc/proto"
	"log"
	"net"
	"testing"
)

func TestHello(t *testing.T) {
	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	pb.RegisterGreeterServer(server, &HelloServer{})

	// 监听端口
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 启动grpc服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
