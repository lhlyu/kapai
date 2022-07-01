package game_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// GetAll 获取所有游戏设定
func (d *dao) GetAll() []*entity.GameEntity {
	cur, err := db.GameCollection.Find(d.ctx, bson.M{})
	if err != nil {
		util.Log.Error().Err(err).Msg("GetAll")
		return nil
	}
	datas := make([]*entity.GameEntity, 0)
	if err := cur.All(d.ctx, &datas); err != nil {
		util.Log.Error().Err(err).Msg("GetAll All")
		return nil
	}
	return datas
}
