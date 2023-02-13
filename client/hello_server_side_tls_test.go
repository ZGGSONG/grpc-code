package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc/proto"
	"log"
	"testing"
)

func TestHelloServerSideTlsClient(t *testing.T) {
	// 根据客户端输入的证书文件和密钥构造 TLS 凭证。
	// 第二个参数 serverNameOverride 为服务名称。
	c, err := credentials.NewClientTLSFromFile("../conf/server-side/server.pem",
		"zggsong.com")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	// 返回一个配置连接的 DialOption 选项。
	// 用于 grpc.Dial(target string, opts ...DialOption) 设置连接选项
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.Message)
}
