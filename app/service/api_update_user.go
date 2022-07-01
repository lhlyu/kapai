package service

import (
	"kapai/app/model"
	"kapai/app/result"
)

// UpdateUser 更新用户信息
func (s *Service) UpdateUser(param *model.UpdateUserParam) result.Result {
	rs := result.NewResult()
	data := s.userDao.GetUser(s.account)
	if data == nil {
		return rs.WithMsg("用户不存在")
	}
	if param.Name != nil {
		data.Name = *param.Name
	}
	if param.Avatar != nil {
		data.Avatar = *param.Avatar
	}
	if param.Password != nil {
		data.Password = *param.Password
	}
	if param.Decks != nil {
		data.Decks = *param.Decks
	}
	if param.Spells != nil {
		data.Spells = *param.Spells
	}
	ok := s.userDao.UpdateUser(data)
	return rs.WithData(ok)
}
