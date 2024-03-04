package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_demo/hello-server/proto"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("服务端被 %v 调用", req.RequestName)
	return &pb.HelloResponse{ResponseReply: "hello" + req.RequestName}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建gRPC服务
	grpcServer := grpc.NewServer()
	// 在rpc服务端中注册我们自己编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
