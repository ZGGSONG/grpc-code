package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc/proto"
	"log"
	"net"
	"testing"
)

func TestHelloServerSideTlsServer(t *testing.T) {
	// 根据服务端输入的证书文件和密钥构造 TLS 凭证
	creds, err := credentials.NewServerTLSFromFile("/Users/song/projects/go/grpc/conf/server-side/server.pem",
		"/Users/song/projects/go/grpc/conf/server-side/server.key")
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	// 返回一个 ServerOption，用于设置服务器连接的凭据。
	// 用于 grpc.NewServer(opt ...ServerOption) 为 gRPC Server 设置连接选项
	lis, err := net.Listen("tcp", "127.0.0.1:50051") //创建 Listen，监听 TCP 端口
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterGreeterServer(s, &HelloServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
