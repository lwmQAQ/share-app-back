package main

import (
	"sync"
	v1 "user-server/cmd/v1"
	"user-server/internal/svc"
)

func main() {
	// 创建服务上下文
	svc := svc.NewServerContext()

	// 使用 WaitGroup 来等待两个服务的启动
	var wg sync.WaitGroup

	// 启动 HTTP 服务
	wg.Add(1)
	go func() {
		defer wg.Done()
		v1.ServerStart(svc)
	}()

	// 启动 RPC 服务
	wg.Add(1)
	go func() {
		defer wg.Done()
		v1.StartRpcServer(svc)
	}()

	// 等待所有服务启动完毕
	wg.Wait()

	// 如果需要阻塞主程序，使用 select {}
	select {}
}
