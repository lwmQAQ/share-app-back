package jiface

/*
解决tcp沾包问题的解决方案，使用设置请求头的方案 将数据使用 |请求头|data格式| 请求头定长 有两个标识 |消息长度|处理方法code| 进行封装
*/
type IDataPack interface {
	Encode(IMessage) ([]byte, error)          // 消息编码
	DecodeHead(head []byte) (IMessage, error) // 消息解码
	GetHeadlen() int                          //获取对应编码方式头部长度

}
