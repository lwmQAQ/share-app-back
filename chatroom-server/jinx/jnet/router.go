package jnet

import "chatroom-server/jinx/jiface"

/*用户可以自定义路由集成 BaseRouter 重写方法 抽象类*/
type BaseRouter struct {
}

//前置钩子
func (b *BaseRouter) PreHandle(request jiface.IRequest) {

}

//业务方法
func (b *BaseRouter) Handle(request jiface.IRequest) {

}

//后置钩子
func (b *BaseRouter) PostHandle(request jiface.IRequest) {

}
