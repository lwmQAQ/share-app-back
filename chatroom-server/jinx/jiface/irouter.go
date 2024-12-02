package jiface

/*
抽象路由接口

*/
type IRouter interface {
	//处理业务之前的钩子方法
	PreHandle(request IRequest)
	//处理业务方法
	Handle(request IRequest)
	//处理业务之后的钩子方法
	PostHandle(request IRequest)
}
