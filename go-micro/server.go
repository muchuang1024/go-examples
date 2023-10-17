package main

import (
	"context"
	"log"

	pb "github.com/caijinlin/golang-notes/go-micro/proto/greeter" // 请替换为你的proto文件路径

	"go-micro.dev/v4" // 引入 go-micro 框架
)

type GreetingService struct{}

func (s *GreetingService) SayHello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloResponse) error {
	rsp.Message = "Hello, " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(GreetingService))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
