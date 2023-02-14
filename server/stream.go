package server

import (
	"fmt"
	pb "grpc/proto/stream"
	"io"
	"log"
	"sync"
	"time"
)

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	// 具体返回多少个response根据业务逻辑调整
	for n := 0; n <= 10; n++ {
		// 通过 send 方法不断推送数据
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 500)
	}
	// 返回nil表示已经完成响应
	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	// for循环接收客户端发送的消息
	for {
		// 通过 Recv() 不断获取客户端 send()推送的消息
		r, err := stream.Recv()
		// err == io.EOF表示已经获取全部数据
		if err == io.EOF {
			// SendAndClose 返回并关闭连接
			// 在客户端发送完毕后服务端即可返回响应
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
		time.Sleep(time.Second)
	}
	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	var (
		wg    sync.WaitGroup //任务编排
		msgCh = make(chan *pb.StreamPoint)
	)
	wg.Add(1)
	go func() {
		n := 0
		defer wg.Done()
		for v := range msgCh {
			err := stream.Send(&pb.StreamResponse{
				Pt: &pb.StreamPoint{
					Name:  v.GetName(),
					Value: int32(n),
				},
			})
			if err != nil {
				fmt.Println("Send error :", err)
				continue
			}
			n++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("recv error :%v", err)
			}
			log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
			msgCh <- &pb.StreamPoint{
				Name: "gRPC Stream Server: Route",
			}
		}
		close(msgCh)
	}()

	wg.Wait() //等待任务结束

	return nil
}
