package server

import (
	"chatroom-server/internal/types"
	"fmt"
	"math/rand"
	"sync"
)

type RoomServer struct {
	Rooms map[string]*Room
	mu    sync.Mutex // 用于保护 Rooms
	r     *rand.Rand // 复用随机数生成器
}

type Room struct {
	RoomID      string   //房间id
	RoomName    string   //房间名称
	IsPrivate   bool     //私密房间
	Password    string   //如果是私密房间要有密码
	RoomCreator uint32   //房主
	RoomMembers []uint32 //成员ID
}

func (s *RoomServer) CreateRoom(userId uint32, req *types.CreateRoomReq) (*Room, error) {
	roomId := s.generateRoomId()
	room := &Room{
		RoomID:      roomId,
		IsPrivate:   req.IsPrivate,
		Password:    req.Password,
		RoomName:    req.RoomName,
		RoomCreator: userId,
		RoomMembers: []uint32{userId}, // 初始化成员列表，包含房主,
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	// 将新房间添加到 Rooms
	s.Rooms[roomId] = room
	return room, nil
}

func (s *RoomServer) generateRoomId() string {
	for {
		roomId := s.generateRoomCode()
		s.mu.Lock() //写锁
		_, exists := s.Rooms[roomId]
		s.mu.Unlock()
		if !exists {
			return roomId
		}
	}
}

func (s *RoomServer) generateRoomCode() string {
	// 使用当前时间作为种子创建本地随机生成器
	roomCode := s.r.Intn(90000000) + 10000000 // 改为8位数
	return fmt.Sprintf("%08d", roomCode)
}
