package client

import (
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc/proto"
	"log"
	"os"
	"testing"
	"time"
)

func TestTlsHelloClient(t *testing.T) {
	// 证书
	certificate, err := tls.LoadX509KeyPair("/Users/song/projects/go/grpc/tls/client.crt",
		"/Users/song/projects/go/grpc/tls/client.key")
	if err != nil {
		panic(err)
		return
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("/Users/song/projects/go/grpc/tls/ca.crt")
	if err != nil {
		panic(err)
		return
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return
	}

	cred := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{certificate},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	})

	// 连接grpc服务器
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	d := pb.NewGreeterClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	resp, err := d.SayHello(ctx, &pb.HelloRequest{Name: "client tls"})
	if err != nil {
		log.Fatalf("could not get info: %v", err)
	}
	log.Printf("GetInfo string: %s", resp.Message)
}
