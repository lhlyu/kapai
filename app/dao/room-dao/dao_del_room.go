package room_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/util"
)

// DelRoom 删除房间
func (d *dao) DelRoom(roomId string) bool {
	if roomId == "" {
		return false
	}
	filter := bson.M{
		"id": roomId,
	}
	_, err := db.UserCollection.DeleteOne(d.ctx, filter)
	if err != nil {
		util.Log.Error().Err(err).Str("roomId", roomId).Msg("DelRoom")
		return false
	}
	return true
}
