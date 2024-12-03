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

type RoomAddRouter struct {
	RoomServer *roomserver.RoomServer
	jnet.BaseRouter
}

func (r *RoomAddRouter) Handle(req jiface.IRequest) {
	fmt.Println("消息ID", req.GetMsgID())
	fmt.Println("收到消息内容是", string(req.GetData()))
	msg := &types.AddRoomReq{}
	err := json.Unmarshal(req.GetData(), msg)
	if err != nil {
		fmt.Println("消息解析出错")
		return
	}
	err = r.RoomServer.AddRoom(req.GetConnection(), msg)
	if err != nil {
		err = req.GetConnection().SendMsg(uint32(message.AddRoomError), []byte("error"))
		if err != nil {
			fmt.Println("写消息错误")
		}
	}

	err = req.GetConnection().SendMsg(uint32(message.AddRoomSuccess), []byte("success"))
	if err != nil {
		fmt.Println("写消息错误")
	}
}
