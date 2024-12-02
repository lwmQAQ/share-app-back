package types

type CreateRoomReq struct {
	IsPrivate bool   `json:"isPrivate"`
	Password  string `json:"password"`
	RoomName  string `json:"roomName"`
}

type AddRoomReq struct {
	RoomID   string `json:"roomId"`
	Password string `json:"password"`
}
