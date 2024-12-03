package rpcpool

import (
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// RpcTaskPool 是一个RPC服务连接池
type RpcTaskPool struct {
	sync.Pool
	connections []*grpc.ClientConn // 存储所有连接
	mu          sync.Mutex         // 互斥锁，保护 connections
}

// GetPool 创建新的 RpcTaskPool
func GetPool(target string, opt ...grpc.DialOption) (*RpcTaskPool, error) {
	pool := &RpcTaskPool{
		Pool: sync.Pool{
			New: func() any {
				conn, err := grpc.Dial(target, opt...)
				if err != nil {
					log.Fatalf("创建连接失败: %v", err)
				}
				return conn
			},
		},
	}
	return pool, nil
}

// Get 获取一个连接
func (c *RpcTaskPool) Get() *grpc.ClientConn {
	conn := c.Pool.Get().(*grpc.ClientConn)
	if conn == nil || conn.GetState() == connectivity.Shutdown || conn.GetState() == connectivity.TransientFailure {
		if conn != nil {
			conn.Close()
		}
		conn = c.Pool.New().(*grpc.ClientConn)
	}
	c.mu.Lock()
	c.connections = append(c.connections, conn) // 记录连接
	c.mu.Unlock()
	return conn
}

// Put 将连接放回连接池
func (c *RpcTaskPool) Put(conn *grpc.ClientConn) {
	if conn == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if conn.GetState() == connectivity.TransientFailure || conn.GetState() == connectivity.Shutdown {
		conn.Close()
	} else {
		c.Pool.Put(conn)
	}
}

// Close 关闭连接池中的所有连接
func (c *RpcTaskPool) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, conn := range c.connections {
		if conn != nil && conn.GetState() != connectivity.Shutdown {
			conn.Close() // 关闭连接
		}
	}
	c.connections = nil // 清空连接列表
}
