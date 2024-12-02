package jiface

type IServer interface {
	Start()
	Stop()
	Sever()
	GetConnManager() IConnManager
	//注册Router方法
	AddRouter(msgId uint32, router IRouter)
}
