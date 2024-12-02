package types

type CreateRoomReq struct {
	IsPrivate bool
	Password  string
	RoomName  string
}
