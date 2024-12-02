package jnet

import (
	"chatroom-server/jinx/jiface"
	"chatroom-server/jinx/utils"
	"fmt"
	"log"
	"net"
	"os"
)

type Server struct {
	Name        string              //服务名称
	IPVersion   string              //协议版本
	IP          string              //监听IP地址
	Port        int                 //监听IP端口
	MsgHandler  jiface.IHandler     //路由
	ConnManager jiface.IConnManager //连接管理器
}

func (s *Server) Start() {

	//获取一个tcp的addr
	addr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, addr)
	if err != nil {
		log.Fatalf("解析地址失败: %v\n", err)
		os.Exit(1)
	}

	// 使用解析的地址启动 TCP 监听器
	listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
	if err != nil {
		log.Fatalf("启动 TCP 监听失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Server is start on", addr)

	//开启工作池
	s.MsgHandler.StartWorkerPool()

	cid := uint32(0)
	//阻塞等待连接
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		//如果连接达到服务器上限
		if s.ConnManager.GetConnNum() >= utils.MyApplication.Server.MaxConn {
			//TODO 给用户返回一个服务器繁忙的数据包
			conn.Close()
			fmt.Println("已经达到连接最大值")
			continue
		}
		dealConn := NewConnection(
			WithConn(conn),
			WithConnID(cid),
			WithServer(s),
			AddHandler(s.MsgHandler),
		)

		go dealConn.Start()
		cid++

	}

}

func (s *Server) GetConnManager() jiface.IConnManager {
	return s.ConnManager
}

func (s *Server) Stop() {
	fmt.Println("尝试关闭服务器 关闭所有连接")
	s.ConnManager.Clear()
	s.MsgHandler.StopWorkerPool()
}

func DefaultRouter() {

}

func (s *Server) Sever() {
	//启动服务
	go s.Start()
	// TODO 这里可以添加启动服务后的操作

	//阻塞主进程
	select {}
}

func (s *Server) AddRouter(msgId uint32, router jiface.IRouter) {
	s.MsgHandler.BindRouter(msgId, router)
	fmt.Println("添加路由成功")
}

/*===================================初始化Server======================================*/
// Option 是一个函数类型，用于配置 Server
type ServerOption func(*Server)

func NewServer(opts ...ServerOption) jiface.IServer {
	//默认选项
	server := &Server{
		Name:        "tcp-server",
		IPVersion:   "tcp4",
		IP:          "0.0.0.0",
		Port:        3000,
		MsgHandler:  NewHandler(),
		ConnManager: NewConnManager(),
	}
	server.IP = utils.MyApplication.Server.Host
	server.Port = utils.MyApplication.Server.Port
	server.Name = utils.MyApplication.Server.Name
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// WithHost 设置 Server 的 Host
func WithHost(ip string) ServerOption {
	return func(s *Server) {
		s.IP = ip
	}
}

// WithPort 设置 Server 的 Port
func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.Port = port
	}
}

// WithPort 设置 Server 的 IP版本
func WithIPVersion(version string) ServerOption {
	return func(s *Server) {
		s.IPVersion = version
	}
}

// WithPort 设置 Server 的 IP版本
func WithName(name string) ServerOption {
	return func(s *Server) {
		s.Name = name
	}
}
