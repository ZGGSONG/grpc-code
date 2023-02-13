package server

import (
	"context"
	pb "grpc/proto"
)

type HelloServer struct {
	pb.UnimplementedGreeterServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "say hello" + req.Name}, nil
}
