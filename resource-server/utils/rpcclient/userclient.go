package rpcclient

import (
	"context"
	"fmt"
	"log"
	"resource-server/internal/rpcclient/userclient"
	"resource-server/utils/etcd"
	"resource-server/utils/rpcpool"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
连接的生命周期： 在 NewRpcClient 函数中，你在成功连接后立即关闭了连接（defer conn.Close()）。
这意味着返回的 RpcClient 结构体中的 RiverClient 将指向一个已经关闭的连接。这是为什么在你后续使用 rpcclient.RiverClient.RunModel(ctx, params) 时会失败。
解决方案： 确保 conn 的生命周期与 RpcClient 的生命周期相同。你可以创建一个结构体来保存连接和客户端，像这样：
*/
type UserRpcClient struct {
	Servers map[string]*rpcpool.RpcTaskPool
}

func NewUserRpcClient(etcdutil *etcd.ETCDUtil) *UserRpcClient {
	server := &UserRpcClient{
		Servers: make(map[string]*rpcpool.RpcTaskPool),
	}
	addrs := etcdutil.DiscoverServices("UserServer", server.UpdateServer)
	fmt.Println(addrs)
	return server
}

func (client *UserRpcClient) GetUserInfo(ctx context.Context, param *userclient.GetUserInfoReq, addr string) (*userclient.GetUserInfoResp, error) {
	conn := client.Servers[addr].Get()
	defer client.Servers[addr].Put(conn)

	if conn == nil {
		return nil, fmt.Errorf("failed to get connection for address: %s", addr)
	}
	userclient := userclient.NewUserServiceClient(conn)
	result, err := userclient.GetUserInfo(ctx, param)
	return result, err
}

/*
在 UpdateServer 方法中，可以遍历传入的新地址列表 (addrs)，并与当前已有的连接池对比：
对于 addrs 中的每个地址，如果当前 Servers 中没有这个地址，则创建一个新的连接池并添加到 Servers。
对于当前 Servers 中的地址，如果它不在 addrs 列表中，说明该服务地址已无效，应关闭对应的连接池并将其从 Servers 中删除。
*/
func (client *UserRpcClient) UpdateServer(addrs []string) {
	fmt.Println("服务发生变动")
	newAddrSet := make(map[string]struct{})
	for _, addr := range addrs {
		newAddrSet[addr] = struct{}{}
	}

	// 添加新的地址
	for _, addr := range addrs {
		if _, exists := client.Servers[addr]; !exists {
			// 如果地址不存在于当前连接池中，创建并添加新连接池
			pool, err := rpcpool.GetPool(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("Failed to connect to %s: %v", addr, err)
				continue
			}
			client.Servers[addr] = pool
			log.Printf("Added new server connection pool for %s", addr)
		}
	}

	// 删除无效的旧地址
	for addr, pool := range client.Servers {
		if _, exists := newAddrSet[addr]; !exists {
			// 如果当前地址不在新地址列表中，关闭连接池并删除
			pool.Close() // 假设连接池具有 Close 方法
			delete(client.Servers, addr)
			log.Printf("Removed obsolete server connection pool for %s", addr)
		}
	}
}

func (client *UserRpcClient) Close() {
	for _, pool := range client.Servers {
		// addr 是键，pool 是值
		pool.Close()
	}
}
