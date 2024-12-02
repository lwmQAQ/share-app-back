package jnet

import (
	"chatroom-server/jinx/jiface"
	"fmt"
	"sync"
)

/*
连接管理模块
*/
type ConnManager struct {
	connections    map[uint32]jiface.IConnection //连接集合
	connLock       sync.RWMutex                  //读写锁
	preConnHandle  func(jiface.IConnection)
	postConnHandle func(jiface.IConnection)
}

// 默认不处理
func Default(jiface.IConnection) {}

func NewConnManager() jiface.IConnManager {
	return &ConnManager{
		connections:    make(map[uint32]jiface.IConnection),
		preConnHandle:  Default,
		postConnHandle: Default,
	}
}

// 处理连接钩子用户可以注册
func (m *ConnManager) BindPreConnection(Func func(jiface.IConnection)) {
	m.preConnHandle = Func
}

// 处理连接删除钩子用户可以注册
func (m *ConnManager) BindPostConnection(Func func(jiface.IConnection)) {
	m.postConnHandle = Func
}

// 添加连接
func (m *ConnManager) AddConnection(conn jiface.IConnection) {
	//用户连接建立钩子
	m.preConnHandle(conn)
	//加写锁
	m.connLock.Lock()
	defer m.connLock.Unlock()
	m.connections[conn.GetConnID()] = conn
	fmt.Println("添加一个连接 ID：", conn.GetConnID())
}

// 删除连接
func (m *ConnManager) DeleteConnection(conn jiface.IConnection) {
	//用户连接断开钩子
	m.postConnHandle(conn)
	m.connLock.Lock()
	defer m.connLock.Unlock()
	delete(m.connections, conn.GetConnID())
	fmt.Println("删除一个连接 ID：", conn.GetConnID())
}

// 根据连接ID获取连接
func (m *ConnManager) GetConnByID(id uint32) (jiface.IConnection, error) {
	//加读锁
	m.connLock.RLock()
	defer m.connLock.Unlock()
	if conn, ok := m.connections[id]; ok {
		return conn, nil
	}
	return nil, fmt.Errorf("没有该连接")
}

// 获取当前连接数量
func (m *ConnManager) GetConnNum() int {
	return len(m.connections)
}

// 清理全部连接
func (m *ConnManager) Clear() {
	m.connLock.Lock()
	defer m.connLock.Unlock()
	for connID, conn := range m.connections {
		//停止所有连接
		conn.Stop()
		delete(m.connections, connID)
	}
	fmt.Println("清除所有连接")
}
