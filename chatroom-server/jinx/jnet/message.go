package jnet

import "chatroom-server/jinx/jiface"

type Message struct {
	Id      uint32 // 消息id
	DataLen uint32 //消息长度
	Data    []byte //消息内容
}

func NewMessage(msgId uint32, data []byte) jiface.IMessage {
	return &Message{
		Id:      msgId,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// 实现 GetMessageId 方法
func (m *Message) GetMessageId() uint32 {
	return m.Id
}

// 实现 GetMessageLen 方法
func (m *Message) GetMessageLen() uint32 {
	return m.DataLen
}

// 实现 GetMessageData 方法
func (m *Message) GetMessageData() []byte {
	return m.Data
}

func (m *Message) SetMessageLen(len uint32) {
	m.DataLen = len
}

// 实现 SetMessageId 方法
func (m *Message) SetMessageId(id uint32) {
	m.Id = id
}

// 实现 SetMessageData 方法
func (m *Message) SetMessageData(data []byte) {
	m.Data = data
	m.DataLen = uint32(len(data)) // 自动更新数据长度
}
