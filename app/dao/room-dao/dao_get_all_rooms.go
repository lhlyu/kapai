package room_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// GetAllRooms 获取所有房间
func (d *dao) GetAllRooms() []*entity.RoomEntity {
	filter := bson.M{}
	cur, err := db.UserCollection.Find(d.ctx, filter)
	if err != nil {
		util.Log.Error().Err(err).Msg("GetAllRooms")
		return nil
	}

	datas := make([]*entity.RoomEntity, 0)
	if err := cur.All(d.ctx, &datas); err != nil {
		util.Log.Error().Err(err).Msg("GetAllRooms All")
		return nil
	}
	return datas
}
