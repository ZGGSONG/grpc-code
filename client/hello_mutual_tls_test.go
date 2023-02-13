package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc/proto"
	"log"
	"os"
	"testing"
)

func TestHelloMutualTlsClient(t *testing.T) {
	// 公钥中读取和解析公钥/私钥对
	cert, err := tls.LoadX509KeyPair("../conf/mutual/client.crt", "../conf/mutual/client.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair error ", err)
		return
	}
	// 创建一组根证书
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("../conf/mutual/ca.crt")
	if err != nil {
		fmt.Println("ReadFile ca.crt error ", err)
		return
	}
	// 解析证书
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		fmt.Println("certPool.AppendCertsFromPEM error ")
		return
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
		//InsecureSkipVerify: true,
		ServerName: "zggsong.com",
	})

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
