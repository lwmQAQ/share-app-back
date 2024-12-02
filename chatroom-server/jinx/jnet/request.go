package jnet

import "chatroom-server/jinx/jiface"

/*
请求模块 封装客户端每一次的请求 将连接和数据封装
*/
type Request struct {
	conn jiface.IConnection
	msg  jiface.IMessage
}

func NewRequest(conn jiface.IConnection, msg jiface.IMessage) jiface.IRequest {
	return &Request{
		conn: conn,
		msg:  msg,
	}
}

func (r *Request) GetConnection() jiface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMessageData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMessageId()
}
