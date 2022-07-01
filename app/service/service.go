package service

import (
	"context"
	room_dao "kapai/app/dao/room-dao"
	user_dao "kapai/app/dao/user-dao"
)

type Service struct {
	ctx     context.Context
	userDao user_dao.Dao
	roomDao room_dao.Dao
	account string
}

func NewService(ctx context.Context, account string) *Service {
	return &Service{
		ctx:     ctx,
		userDao: user_dao.NewDao(ctx),
		roomDao: room_dao.NewDao(ctx),
		account: account,
	}
}
