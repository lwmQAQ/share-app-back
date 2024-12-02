package jiface

/*
handler 模块实现多路由
*/
type IHandler interface {
	BindRouter(uint32, IRouter)
	UseRouter(IRequest)
	StartWorkerPool()
	StopWorkerPool()
	HandleRequest(IRequest)
}
