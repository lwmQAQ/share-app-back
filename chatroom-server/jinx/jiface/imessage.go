package jiface

/*
消息封装模块
*/
type IMessage interface {
	GetMessageId() uint32   //获取消息id
	GetMessageLen() uint32  //获取消息长度
	GetMessageData() []byte //获取消息内容
	SetMessageId(uint32)
	SetMessageLen(uint32)
	SetMessageData([]byte)
}
