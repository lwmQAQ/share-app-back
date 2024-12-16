package v1

import (
	"fmt"
	"user-server/internal/rpcclient/userserver"
	"user-server/internal/server/rpcuserver"
	"user-server/internal/svc"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartRpcServer(svc *svc.ServiceContext) {
	// 启动 gRPC 服务器
	addr := fmt.Sprintf("%s:%d", svc.ServerConfig.RpcServer.Host, svc.ServerConfig.RpcServer.Port)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		svc.Logger.Errorf("端口监听失败 %v", err)
		return
	}

	s := grpc.NewServer()
	userserver.RegisterUserServiceServer(s, rpcuserver.NewRpcUserServer(svc))

	// 注册反射服务（可选，用于 gRPC CLI 调试）
	reflection.Register(s)
	svc.EtcdUtil.RegisterService(svc.ServerConfig.RpcServer.ServerName, addr, 5)
	fmt.Println("Server is listening on port 50051...")
	// 启动服务器
	if err := s.Serve(listen); err != nil {
		svc.Logger.Errorf("服务启动失败 %v", err)
	}

}
