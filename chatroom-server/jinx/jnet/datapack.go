package jnet

import (
	"chatroom-server/jinx/jiface"
	"chatroom-server/jinx/utils"
	"encoding/binary"
	"fmt"
)

/*
解决tcp沾包问题的解决方案，使用设置请求头的方案 将数据使用 |请求头|data格式| 请求头定长 有两个标识 |消息长度|处理方法code| 进行封装
*/

type DataPack struct {
	headlen int
}

func NewDataPack() jiface.IDataPack {
	return &DataPack{
		headlen: 8,
	}
}

func (dp *DataPack) GetHeadlen() int {
	return dp.headlen
}

func (dp *DataPack) Encode(msg jiface.IMessage) ([]byte, error) { // 消息编码
	data := msg.GetMessageData()
	msgLen := uint32(len(data))                                // 消息体长度
	header := make([]byte, 8)                                  // 请求头长度固定 8 字节
	binary.BigEndian.PutUint32(header[:4], msgLen)             // 消息长度
	binary.BigEndian.PutUint32(header[4:], msg.GetMessageId()) // 方法代码

	return append(header, data...), nil
}

func (dp *DataPack) DecodeHead(head []byte) (jiface.IMessage, error) { // 消息解码
	len := len(head)
	if len < 8 {
		return nil, fmt.Errorf("data too short to read header")
	} else if len > int(utils.MyApplication.Server.MaxPackageSize) {
		return nil, fmt.Errorf("data too Long")
	}

	msg := &Message{}
	// 解析请求头
	msgLen := binary.BigEndian.Uint32(head[:4])
	methodCode := binary.BigEndian.Uint32(head[4:8])
	msg.SetMessageId(methodCode)
	msg.SetMessageLen(msgLen)
	return msg, nil

}
