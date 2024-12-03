package v1

import (
	"chatroom-server/internal/handle"
	"chatroom-server/internal/message"
	roomserver "chatroom-server/internal/server"
	"chatroom-server/internal/svc"
	"chatroom-server/jinx/jnet"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func ServerStart() {
	server := jnet.NewServer()
	ctx := svc.NewServerContext()
	Roomserver := roomserver.NewRoomServer(ctx)
	server.AddRouter(uint32(message.CreateRoom), &handle.RoomCreateRouter{
		RoomServer: Roomserver,
	})
	server.AddRouter(uint32(message.AddRoom), &handle.RoomAddRouter{
		RoomServer: Roomserver,
	})
	server.AddRouter(uint32(message.SetUserAttribute), &handle.UserAttributeRouter{})

	// 使用 goroutine 启动服务器
	go func() {
		server.Start()
	}()

	// 捕获 Ctrl+C 或终止信号
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// 阻塞等待信号
	<-stopChan
	fmt.Println("\n收到终止信号，正在关闭服务器...")

	// 优雅关闭服务器
	server.Stop()

	fmt.Println("Server stopped.")

}
