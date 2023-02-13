package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/proto"
	"log"
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	// 建立链接
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	// 退出时关闭链接
	defer conn.Close()

	// 实例化客户端
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 发起请求
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "testHello"})
	if err != nil {
		log.Fatalf("cloud not get info, err: %v", err)
	}
	log.Printf("%s", resp.Message)
}
