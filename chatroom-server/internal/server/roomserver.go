package roomserver

import (
	"chatroom-server/internal/types"
	"chatroom-server/jinx/jiface"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RoomServer struct {
	Rooms map[string]*Room
	mu    sync.Mutex // 用于保护 Rooms
	r     *rand.Rand // 复用随机数生成器
}

type Room struct {
	RoomID      string               //房间id
	RoomName    string               //房间名称
	IsPrivate   bool                 //私密房间
	Password    string               //如果是私密房间要有密码
	RoomCreator uint32               //房主
	RoomMembers []jiface.IConnection //成员ID
}

func NewRoomServer() *RoomServer {
	return &RoomServer{
		Rooms: make(map[string]*Room),                      // 初始化 Rooms 映射
		r:     rand.New(rand.NewSource(time.Now().Unix())), // 初始化随机数生成器
	}
}

func (s *RoomServer) CreateRoom(user jiface.IConnection, req *types.CreateRoomReq) (*types.CreateRoomResp, error) {
	roomId := s.generateRoomId()
	room := &Room{
		RoomID:      roomId,
		IsPrivate:   req.IsPrivate,
		Password:    req.Password,
		RoomName:    req.RoomName,
		RoomCreator: user.GetConnID(),
		RoomMembers: []jiface.IConnection{user}, // 初始化成员列表，包含房主,
	}
	s.mu.Lock()
	// 将新房间添加到 Rooms
	s.Rooms[roomId] = room
	s.mu.Unlock()
	//TODO 通过rpc获取用户具体信息
	resp := &types.CreateRoomResp{
		RoomID:     roomId,
		RoomName:   room.RoomName,
		MembersNum: len(room.RoomMembers),
		IsPrivate:  room.IsPrivate,
		Members:    nil,
	}
	return resp, nil
}

func (s *RoomServer) AddRoom(user jiface.IConnection, req *types.AddRoomReq) error {
	fmt.Println("1")
	if room, ok := s.Rooms[req.RoomID]; ok {
		if !room.IsPrivate || req.Password == room.Password {
			room.RoomMembers = append(room.RoomMembers, user)
			fmt.Println("yes")
			return nil
		}
		return fmt.Errorf("密码错误")
	}
	return fmt.Errorf("房间不存在")
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
