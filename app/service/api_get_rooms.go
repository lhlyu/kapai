package service

import (
	"kapai/app/model"
	"kapai/app/result"
)

// GetRooms 获取所有的房间
func (s *Service) GetRooms() result.Result[[]*model.Room] {
	rs := result.NewResult[[]*model.Room]()
	datas := s.roomDao.GetAllRooms()
	if len(datas) == 0 {
		return rs
	}

	accounts := make([]string, 0)

	for i := range datas {
		if datas[i].Black != "" {
			accounts = append(accounts, datas[i].Black)
		}
		if datas[i].Red != "" {
			accounts = append(accounts, datas[i].Red)
		}
	}

	// 获取账号信息
	userMap := s.userDao.GetUsers(accounts)

	list := make([]*model.Room, 0)
	for i := range datas {
		r := &model.Room{
			Id:     datas[i].Id,
			Status: datas[i].Status,
		}
		if v, ok := userMap[datas[i].Black]; ok {
			r.Black = &model.UserModel{
				Account: v.Account,
				Name:    v.Name,
				Avatar:  v.Avatar,
			}
		}
		if v, ok := userMap[datas[i].Red]; ok {
			r.Red = &model.UserModel{
				Account: v.Account,
				Name:    v.Name,
				Avatar:  v.Avatar,
			}
		}
		list = append(list, r)
	}
	return rs.WithData(list)
}
