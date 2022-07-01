package service

import (
	"kapai/app/entity"
	"kapai/app/model"
	"kapai/app/result"
	"kapai/app/util"
)

// LeaveRoom 离开房间
func (s *Service) LeaveRoom(param *model.RoomParam) result.Result[bool] {
	rs := result.NewResult[bool]()
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

	if room.Black == s.account {
		room.Black = ""
	}
	if room.Red == s.account {
		room.Red = ""
	}

	user.RoomId = ""
	s.userDao.UpdateUser(user)

	if room.Black == "" && room.Red == "" {
		ok := s.roomDao.DelRoom(room.Id)
		return rs.WithData(ok)
	}

	ok := s.roomDao.UpdateRoom(room)
	return rs.WithData(ok)
}
