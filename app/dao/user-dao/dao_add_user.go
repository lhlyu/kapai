package user_dao

import (
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
)

// AddUser 添加用户
func (d *dao) AddUser(userEntity *entity.UserEntity) bool {
	if userEntity == nil {
		return false
	}
	_, err := db.UserCollection.InsertOne(d.ctx, userEntity)
	if err != nil {
		util.Log.Error().Err(err).Interface("userEntity", userEntity).Msg("AddUser")
		return false
	}
	return true
}
