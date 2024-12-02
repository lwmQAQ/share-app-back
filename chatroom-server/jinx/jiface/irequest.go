package jiface

/*
请求封装模块
*/
type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMsgID() uint32
}
