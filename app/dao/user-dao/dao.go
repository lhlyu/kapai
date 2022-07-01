package user_dao

import (
	"context"
	"kapai/app/entity"
)

type Dao interface {
	// AddUser 添加用户
	AddUser(userEntity *entity.UserEntity) bool
	// UpdateUser 更新用户信息
	UpdateUser(userEntity *entity.UserEntity) bool
	// GetUser 获取用户信息
	GetUser(account string) *entity.UserEntity
	// GetUsers 批量获取多个用户信息
	GetUsers(accounts []string) map[string]*entity.UserEntity
}

type dao struct {
	ctx context.Context
}

func NewDao(ctx context.Context) Dao {
	return &dao{ctx}
}
