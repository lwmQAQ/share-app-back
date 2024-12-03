package svc

import (
	"chatroom-server/utils/etcd"
	"chatroom-server/utils/rpcclient"
)

type ServiceContext struct {
	EtcdUtil      *etcd.ETCDUtil
	UserRpcClient *rpcclient.UserRpcClient
}

func NewServerContext() *ServiceContext {
	etcdutil := etcd.NewETCDUtil()
	userpclient := rpcclient.NewRpcClient(etcdutil)
	return &ServiceContext{
		EtcdUtil:      etcdutil,
		UserRpcClient: userpclient,
	}
}
