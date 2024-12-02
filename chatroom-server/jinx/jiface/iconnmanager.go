package jiface

/*
连接管理模块
*/

type IConnManager interface {
	//处理连接钩子用户可以注册
	BindPreConnection(Func func(IConnection))
	//处理连接删除钩子用户可以注册
	BindPostConnection(Func func(IConnection))
	//添加连接
	AddConnection(IConnection)
	//删除连接
	DeleteConnection(IConnection)
	//根据连接ID获取连接
	GetConnByID(id uint32) (IConnection, error)
	//获取当前连接数量
	GetConnNum() int
	//清理全部连接
	Clear()
}
