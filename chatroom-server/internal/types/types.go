package types

type CreateRoomReq struct {
	IsPrivate bool   `json:"isPrivate"`
	Password  string `json:"password"`
	RoomName  string `json:"roomName"`
}
type CreateRoomResp struct {
	MembersNum int      `json:"membersNum"`
	IsPrivate  bool     `json:"isPrivate"`
	RoomID     string   `json:"roomId"`
	RoomName   string   `json:"roomName"`
	Members    []Member `json:"members"`
}

type Member struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Sex      int    `json:"sex"`
}
type AddRoomReq struct {
	RoomID   string `json:"roomId"`
	Password string `json:"password"`
}

type SendTextMsg struct {
	RoomID string `json:"roomId"`
	Text   string `json:"text"`
}
