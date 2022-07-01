package room_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// UpdateRoom 更新房间
func (d *dao) UpdateRoom(param *entity.RoomEntity) bool {
	if param.Id == "" {
		return false
	}

	filter := bson.M{
		"id": param.Id,
	}
	update := bson.M{
		"$set": bson.M{
			"status": param.Status,
			"black":  param.Black,
			"red":    param.Red,
		},
	}

	if _, err := db.RoomCollection.UpdateOne(d.ctx, filter, update); err != nil {
		util.Log.Error().Err(err).Interface("param", param).Msg("UpdateRoom")
		return false
	}

	return true
}
