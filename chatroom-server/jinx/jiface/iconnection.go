package jiface

import "net"

/*
客户端连接模块
*/
type IConnection interface {
	//启动连接
	Start()
	//停止连接
	Stop()
	//获取连接的conn
	GetTCPConnection() *net.TCPConn
	//获取当前连接的ID
	GetConnID() uint32
	//获取连接远程端的状态
	RemoteADdr() net.Addr
	//发送数据
	SendMsg(msgId uint32, data []byte) error
	//设置连接属性
	SetProperty(key string, value any)
	//获取连接属性
	GetProperty(key string) (any, error)
	//删除属性
	DeleteProperty(key string) error
}

//当前连接绑定的处理方法
type HandleFunc func(*net.TCPConn, []byte, int) error
