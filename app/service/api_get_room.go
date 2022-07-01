package service

import (
	"kapai/app/model"
	"kapai/app/result"
	"kapai/app/util"
)

// GetRoom 获取单个房间的信息
func (s *Service) GetRoom(roomId string) result.Result {
	rs := result.NewResult()
	data := s.roomDao.GetRoom(roomId)
	if data == nil {
		return rs.WithMsg("房间不存在")
	}

	var (
		black *model.UserModel
		red   *model.UserModel
	)
	util.WaitGroup(
		func() {
			if data.Black == "" {
				return
			}
			user := s.userDao.GetUser(data.Black)
			if user == nil {
				return
			}
			black = &model.UserModel{
				Account: user.Account,
				Name:    user.Name,
				Avatar:  user.Avatar,
			}
		},
		func() {
			if data.Red == "" {
				return
			}
			user := s.userDao.GetUser(data.Red)
			if user == nil {
				return
			}
			red = &model.UserModel{
				Account: user.Account,
				Name:    user.Name,
				Avatar:  user.Avatar,
			}
		},
	)

	return rs.WithData(&model.Room{
		Id:     data.Id,
		Status: data.Status,
		Black:  black,
		Red:    red,
	})
}
