package server

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc/proto"
	"log"
	"net"
	"os"
	"testing"
)

func TestTlsHelloServe(t *testing.T) {
	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 证书
	certificate, err := tls.LoadX509KeyPair("/Users/song/projects/go/grpc/tls/server.crt",
		"/Users/song/projects/go/grpc/tls/server.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("/Users/song/projects/go/grpc/tls/ca.crt")
	if err != nil {
		panic(err)

	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		panic("AppendCertsFromPEM failed")
	}

	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})
	// 通过tls证书实例化grpc服务端
	s := grpc.NewServer(grpc.Creds(cred))

	// 注册Greeter服务
	pb.RegisterGreeterServer(s, &HelloServer{})

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
