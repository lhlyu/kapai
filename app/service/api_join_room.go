package service

import (
	"kapai/app/entity"
	"kapai/app/model"
	"kapai/app/result"
	"kapai/app/util"
)

const (
	ROLE_PLAYER   = "player"
	ROLE_AUDIENCE = "audience"
)

// JoinRoom 加入房间
// 返回：player - 玩家; audience - 观众
func (s *Service) JoinRoom(param *model.RoomParam) result.Result[string] {
	rs := result.NewResult[string]()
	var (
		room *entity.RoomEntity
		user *entity.UserEntity
	)
	util.WaitGroup(
		func() {
			room = s.roomDao.GetRoom(param.RoomId)
		},
		func() {
			user = s.userDao.GetUser(s.account)
		},
	)
	if room == nil {
		return rs.WithMsg("房间不存在")
	}
	if user == nil {
		return rs.WithMsg("用户不存在")
	}
	if room.Black != "" && room.Red != "" {
		return rs.WithData(ROLE_AUDIENCE)
	}
	joined := false
	if room.Black == "" {
		room.Black = s.account
		joined = true
	}
	if !joined && room.Red == "" {
		room.Red = s.account
	}
	if !s.roomDao.UpdateRoom(room) {
		return rs.WithMsg("加入房间异常")
	}
	user.RoomId = room.Id
	s.userDao.UpdateUser(user)
	return rs.WithData(ROLE_PLAYER)
}
