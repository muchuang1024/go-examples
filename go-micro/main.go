// Package main
package main

import (
	"context"
	"log"

	"go-micro.dev/v4" // 引入 go-micro 框架
)

/*
* 定义一个 handler
 */

// handler 接口请求参数
type Request struct {
	Name string `json:"name"`
}

// handler接口请求响应
type Response struct {
	Message string `json:"message"`
}

// Helloworld 服务
type Helloworld struct{}

// 实现 Helloworld 服务接口 Greeting handler
func (h *Helloworld) Greeting(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {

	// 创建微服务实例
	service := micro.NewService(
		micro.Name("helloworld"),      // 名称
		micro.Address(":8080"),        // 端口
		micro.Handle(new(Helloworld)), // 注册 Helloworld 服务
	)

	// 接收命令行参数  如--server_address
	service.Init()

	// 运行微服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
