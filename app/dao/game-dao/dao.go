package game_dao

import (
	"context"
	"kapai/app/entity"
)

type Dao interface {
	// GetAll 获取所有游戏设定
	GetAll() []*entity.GameEntity
}

type dao struct {
	ctx context.Context
}

func NewDao(ctx context.Context) Dao {
	return &dao{ctx}
}
