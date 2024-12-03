package handle

import (
	"chatroom-server/internal/message"
	"chatroom-server/internal/types"
	"chatroom-server/jinx/jiface"
	"chatroom-server/jinx/jnet"
	"encoding/json"
	"fmt"
)

type UserAttributeRouter struct {
	jnet.BaseRouter
}

func (r *UserAttributeRouter) Handle(req jiface.IRequest) {
	fmt.Println("消息ID", req.GetMsgID())
	fmt.Println("收到消息内容是", string(req.GetData()))
	msg := &types.SetUserAttribute{}
	err := json.Unmarshal(req.GetData(), msg)
	if err != nil {
		fmt.Printf("消息解析出错 %v", err)
		err = req.GetConnection().SendMsg(uint32(message.SetUserAttributeError), nil)
		if err != nil {
			fmt.Println("写消息错误")
		}
		return
	}
	req.GetConnection().SetProperty(msg.Key, msg.Value)
	err = req.GetConnection().SendMsg(uint32(message.SendTextMsgError), nil)
	if err != nil {
		fmt.Println("写消息错误")
	}
}
