package room_dao

import (
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// AddRoom 添加房间
func (d *dao) AddRoom(roomEntity *entity.RoomEntity) bool {
	if roomEntity == nil {
		return false
	}
	_, err := db.UserCollection.InsertOne(d.ctx, roomEntity)
	if err != nil {
		util.Log.Error().Err(err).Interface("roomEntity", roomEntity).Msg("AddRoom")
		return false
	}
	return true
}
