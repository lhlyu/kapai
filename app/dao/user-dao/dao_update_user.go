package user_dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"kapai/app/db"
	"kapai/app/entity"
	"kapai/app/util"
	"time"
)

// UpdateUser 更新用户信息
func (d *dao) UpdateUser(userEntity *entity.UserEntity) bool {
	if userEntity == nil || userEntity.Account == "" {
		return false
	}

	filter := bson.M{
		"account": userEntity.Account,
	}

	update := bson.M{
		"$set": bson.M{
			"updateTime": time.Now().UnixMilli(),
			"name":       userEntity.Name,
			"avatar":     userEntity.Avatar,
			"password":   userEntity.Password,
			"decks":      userEntity.Decks,
			"spells":     userEntity.Spells,
			"roomId":     userEntity.RoomId,
		},
	}

	_, err := db.UserCollection.UpdateOne(d.ctx, filter, update)
	if err != nil {
		util.Log.Error().Err(err).Interface("userEntity", userEntity).Msg("UpdateUser")
		return false
	}
	return true
}
