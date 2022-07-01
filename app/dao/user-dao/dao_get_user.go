package user_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// GetUser 获取用户信息
func (d *dao) GetUser(account string) *entity.UserEntity {
	if account == "" {
		return nil
	}
	filter := bson.M{
		"account": account,
	}
	r := db.UserCollection.FindOne(d.ctx, filter)
	if r.Err() != nil {
		util.Log.Error().Err(r.Err()).Str("account", account).Msg("GetUser")
		return nil
	}
	data := &entity.UserEntity{}
	if err := r.Decode(data); err != nil {
		util.Log.Error().Err(r.Err()).Str("account", account).Msg("GetUser Decode")
		return nil
	}
	if data.Account == "" {
		return nil
	}
	return data
}
