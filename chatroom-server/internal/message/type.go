package message

type MessageType int

const (
	CreateRoom MessageType = iota + 1
	CreateRoomSuccess
	CreateRoomError
	AddRoom
	AddRoomSuccess
	AddRoomError
)
