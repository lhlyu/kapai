package room_dao

import (
	"context"
	"kapai/app/entity"
)

type Dao interface {
	// AddRoom 添加房间
	AddRoom(roomEntity *entity.RoomEntity) bool

	// DelRoom 删除房间
	DelRoom(roomId string) bool

	// GetRoom 获取房间信息
	GetRoom(roomId string) *entity.RoomEntity

	// GetAllRooms 获取所有房间
	GetAllRooms() []*entity.RoomEntity

	// UpdateRoom 更新房间
	UpdateRoom(param *entity.RoomEntity) bool
}

type dao struct {
	ctx context.Context
}

func NewDao(ctx context.Context) Dao {
	return &dao{ctx}
}
