package jnet

import (
	"chatroom-server/jinx/jiface"
	"fmt"
	"io"
	"net"
	"sync"
)

/*
连接模块 封装连接对象
*/
type Connection struct {
	//当前的连接
	Conn *net.TCPConn
	//当前连接的ID
	ConnID uint32
	//当前连接状态
	isClose bool
	//告知连接关闭的通道
	ExitChan chan bool
	//消息管道
	MsgChan chan []byte
	//路由
	MsgHandler jiface.IHandler
	//这个连接属于哪个server
	server jiface.IServer
	//连接属性集合
	property map[string]any
	//属性锁
	propertyLock sync.RWMutex
}

/*==================================给客户端写消息=======================================*/
func (c *Connection) StartWriter() {
	fmt.Println("Write Coroutine is running...", c.ConnID)
	defer fmt.Println("Writer down")
	for {
		select {
		case msg := <-c.MsgChan:
			_, err := c.Conn.Write(msg)
			if err != nil {
				fmt.Println("给客户端发送消息失败")
				continue
			}
		case <-c.ExitChan: //退出
			return
		}
	}
}

/*==================================读客户端消息=======================================*/
func (c *Connection) StartReader() {
	fmt.Println("Reader Coroutine is running...", c.ConnID)
	defer func() {
		fmt.Println("Reader down")
		c.Stop()
	}()
	dp := NewDataPack()
	for {
		// 读取客户端数据到缓冲区中
		//将消息解码
		headlen := dp.GetHeadlen()
		buf := make([]byte, headlen)
		_, err := io.ReadFull(c.Conn, buf)
		if err != nil {
			fmt.Println("读取数据失败", err)
			break
		}
		msg, err := dp.DecodeHead(buf)
		if err != nil {
			fmt.Println("解析请求头错误")
			break
		}
		data := make([]byte, msg.GetMessageLen())
		_, err = io.ReadFull(c.Conn, data)
		if err != nil {
			fmt.Println("解析请求数据错误")
			break
		}
		msg.SetMessageData(data)
		//封装Requst 数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		//发送的消息队列中等待处理
		c.MsgHandler.HandleRequest(&req)
	}
}

// 发送消息方法
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClose {
		return fmt.Errorf("conn is close")
	}
	msg := NewMessage(msgId, data)
	dp := NewDataPack()
	sendmsg, err := dp.Encode(msg)
	if err != nil {
		fmt.Println("封装数据出错", err)
		return fmt.Errorf("封装数据出错")
	}
	c.MsgChan <- sendmsg
	return nil
}

// 设置连接属性
func (c *Connection) SetProperty(key string, value any) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if _, ok := c.property[key]; ok {
		fmt.Println(key, "属性被更改")
	}
	c.property[key] = value

}

func (c *Connection) DeleteProperty(key string) error {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if _, ok := c.property[key]; ok {
		delete(c.property, key)
		return nil
	}
	return fmt.Errorf("没有该属性")
}

func (c *Connection) GetProperty(key string) (any, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("没有该属性")
}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnID = ", c.ConnID)
	//启动当前读数据的业务
	go c.StartReader()
	go c.StartWriter()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().... ConnID=", c.ConnID)
	if c.isClose {
		return
	}

	//标识连接关闭
	c.isClose = true

	//连接关闭
	c.Conn.Close()

	c.ExitChan <- true

	//移除连接
	c.server.GetConnManager().DeleteConnection(c)
	//回收资源
	close(c.ExitChan)
	close(c.MsgChan)
}

// 获取连接的conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取连接远程端的状态
func (c *Connection) RemoteADdr() net.Addr {
	return nil
}

/*==========================================初始化Connection================================================*/
type ConnOption func(*Connection)

func NewConnection(opts ...ConnOption) jiface.IConnection {

	connection := &Connection{
		isClose:    false,
		ExitChan:   make(chan bool, 1),
		MsgHandler: NewHandler(),
		MsgChan:    make(chan []byte),
		property:   map[string]any{},
	}
	for _, opt := range opts {
		opt(connection)
	}

	connection.server.GetConnManager().AddConnection(connection)
	return connection
}

func WithConn(conn *net.TCPConn) ConnOption {
	return func(c *Connection) {
		c.Conn = conn
	}
}

func WithConnID(connid uint32) ConnOption {
	return func(c *Connection) {
		c.ConnID = connid
	}
}
func AddHandler(handler jiface.IHandler) ConnOption {
	return func(c *Connection) {
		c.MsgHandler = handler
	}
}

func WithServer(server jiface.IServer) ConnOption {
	return func(c *Connection) {
		c.server = server
	}
}
