package utils

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"user-server/config"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ETCDUtil struct {
	cli      *clientv3.Client
	services map[string]*serviceInfo //对应服务的客户端
	mu       sync.Mutex
}

type serviceInfo struct {
	addresses         []string
	currentIndex      int
	serviceUpdateHook func([]string) // 新增回调函数
}

// NewETCDUtil 创建新的 ETCDUtil 实例
func NewETCDUtil(config *config.EtcdConfig) *ETCDUtil {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Addrs,
		DialTimeout: time.Duration(config.Timeout) * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	return &ETCDUtil{
		cli:      cli,
		services: make(map[string]*serviceInfo),
	}
}

// RegisterService 注册服务
func (etcd *ETCDUtil) RegisterService(serviceKey string, serviceAddr string, ttl int64) {
	go func() {
		leaseResp, err := etcd.cli.Grant(context.Background(), ttl)
		if err != nil {
			log.Fatalf("Failed to create lease: %v", err)
		}

		key := fmt.Sprintf("%s/%s", serviceKey, serviceAddr)
		_, err = etcd.cli.Put(context.Background(), key, serviceAddr, clientv3.WithLease(leaseResp.ID))
		if err != nil {
			log.Fatalf("Failed to put key with lease: %v", err)
		}
		log.Printf("Service registered: %s at %s", serviceKey, serviceAddr)

		ch, err := etcd.cli.KeepAlive(context.Background(), leaseResp.ID)
		if err != nil {
			log.Fatalf("Failed to set up lease keep-alive: %v", err)
		}

		for keepAliveResp := range ch {
			log.Printf("Keep-alive response for %s - lease ID: %v, TTL: %d", serviceKey, keepAliveResp.ID, keepAliveResp.TTL)
		}
		log.Printf("Lease keep-alive channel closed for service: %s", serviceAddr)
	}()
}

// DiscoverServices 监听多个服务的变更
func (etcd *ETCDUtil) DiscoverServices(serviceKey string, serviceUpdateHook func([]string)) []string {
	addrs, err := etcd.GetAllServiceAddresses(serviceKey)
	if err != nil {
		log.Printf("%s服务出现问题", serviceKey)
	}
	info := &serviceInfo{
		currentIndex:      0,
		serviceUpdateHook: serviceUpdateHook,
		addresses:         addrs,
	}
	etcd.services[serviceKey] = info
	go etcd.watchServices(serviceKey)
	etcd.updateServiceList(serviceKey)
	return addrs
}

// watchServices 监听指定服务的变更
func (etcd *ETCDUtil) watchServices(serviceKey string) {
	rch := etcd.cli.Watch(context.Background(), serviceKey, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			log.Printf("Watch event: %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
		etcd.updateServiceList(serviceKey)
	}
}

// updateServiceList 更新服务列表
func (etcd *ETCDUtil) updateServiceList(serviceKey string) {
	resp, err := etcd.cli.Get(context.Background(), serviceKey, clientv3.WithPrefix())
	if err != nil {
		log.Fatalf("Failed to get services: %v", err)
	}

	addresses := []string{}
	for _, kv := range resp.Kvs {
		addresses = append(addresses, string(kv.Value))
	}

	etcd.mu.Lock()
	if _, exists := etcd.services[serviceKey]; !exists {
		etcd.services[serviceKey] = &serviceInfo{}
	}
	//
	etcd.services[serviceKey].addresses = addresses
	//调用初始化时的回调函数
	etcd.services[serviceKey].serviceUpdateHook(addresses)
	etcd.mu.Unlock()

	log.Printf("Updated service list for %s: %v", serviceKey, addresses)
	//通知
}

// GetServiceInstance 获取指定服务的实例（轮询方式）
func (etcd *ETCDUtil) GetServiceInstance(serviceKey string) (string, error) {
	etcd.mu.Lock()
	defer etcd.mu.Unlock()

	service, exists := etcd.services[serviceKey]
	fmt.Println(service)
	if !exists || len(service.addresses) == 0 {
		return "", fmt.Errorf("no available services for %s", serviceKey)
	}

	address := service.addresses[service.currentIndex]
	service.currentIndex = (service.currentIndex + 1) % len(service.addresses)
	return address, nil
}

// Close 关闭etcd连接
func (etcd *ETCDUtil) Close() {
	etcd.cli.Close()
}

// GetAllServiceAddresses 获取所有服务的地址

func (etcd *ETCDUtil) GetAllServiceAddresses(serviceKey string) ([]string, error) {
	etcd.mu.Lock()
	defer etcd.mu.Unlock()

	// 从 ETCD 查询服务地址
	resp, err := etcd.cli.Get(context.Background(), serviceKey, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("failed to get services from etcd: %v", err)
	}

	var addresses []string
	for _, kv := range resp.Kvs {
		addresses = append(addresses, string(kv.Value))
	}
	return addresses, nil
}
