package room_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// GetRoom 获取房间信息
func (d *dao) GetRoom(roomId string) *entity.RoomEntity {
	if roomId == "" {
		return nil
	}
	filter := bson.M{
		"id": roomId,
	}
	r := db.UserCollection.FindOne(d.ctx, filter)
	if r.Err() != nil {
		util.Log.Error().Err(r.Err()).Str("roomId", roomId).Msg("GetRoom")
		return nil
	}
	data := &entity.RoomEntity{}
	if err := r.Decode(data); err != nil {
		util.Log.Error().Err(r.Err()).Str("roomId", roomId).Msg("GetRoom Decode")
		return nil
	}
	if data.Id == "" {
		return nil
	}
	return data
}
