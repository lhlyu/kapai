package user_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// GetUsers 批量获取多个用户信息
func (d *dao) GetUsers(accounts []string) map[string]*entity.UserEntity {
	if len(accounts) == 0 {
		return nil
	}
	filter := bson.M{
		"account": bson.M{
			"$in": accounts,
		},
	}
	cur, err := db.UserCollection.Find(d.ctx, filter)
	if err != nil {
		util.Log.Error().Err(err).Interface("accounts", accounts).Msg("GetUsers")
		return nil
	}

	datas := make([]*entity.UserEntity, 0)
	if err := cur.All(d.ctx, &datas); err != nil {
		util.Log.Error().Err(err).Interface("accounts", accounts).Msg("GetUsers All")
		return nil
	}

	m := make(map[string]*entity.UserEntity)
	for i := range datas {
		if datas[i] == nil || datas[i].Account == "" {
			continue
		}
		m[datas[i].Account] = datas[i]
	}
	return m
}
