package handle

import (
	"chatroom-server/internal/message"
	roomserver "chatroom-server/internal/server"
	"chatroom-server/internal/types"
	"chatroom-server/jinx/jiface"
	"chatroom-server/jinx/jnet"
	"encoding/json"
	"fmt"
)

type RoomCreateRouter struct {
	RoomServer *roomserver.RoomServer
	jnet.BaseRouter
}

func (r *RoomCreateRouter) Handle(req jiface.IRequest) {
	fmt.Println("消息ID", req.GetMsgID())
	fmt.Println("收到消息内容是", string(req.GetData()))
	msg := &types.CreateRoomReq{}
	err := json.Unmarshal(req.GetData(), msg)
	if err != nil {
		fmt.Printf("消息解析出错 %v", err)
		return
	}
	resp, err := r.RoomServer.CreateRoom(req.GetConnection(), msg)
	if err != nil {
		err = req.GetConnection().SendMsg(uint32(message.CreateRoomError), nil)
		if err != nil {
			fmt.Println("写消息错误")
		}
	}

	// 将结构体转换为 JSON 字节数组
	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("解析结构体错误")
	}
	err = req.GetConnection().SendMsg(uint32(message.CreateRoomSuccess), jsonData)
	if err != nil {
		fmt.Println("写消息错误")
	}
}
