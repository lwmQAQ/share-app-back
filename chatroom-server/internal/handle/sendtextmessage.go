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

type SendTextRouter struct {
	RoomServer *roomserver.RoomServer
	jnet.BaseRouter
}

func (r *SendTextRouter) Handle(req jiface.IRequest) {
	msg := &types.SendTextMsg{}
	err := json.Unmarshal(req.GetData(), msg)
	if err != nil {
		fmt.Println("消息解析出错")
		return
	}
	if room, ok := r.RoomServer.Rooms[msg.RoomID]; ok {
		for _, member := range room.RoomMembers {
			err := member.SendMsg(uint32(message.TextMessage), []byte(msg.Text))
			if err != nil {
				fmt.Println("发送消息失败")
			}
		}
	}

}
