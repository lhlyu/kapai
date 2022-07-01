package service

import (
	"kapai/app/entity"
	"kapai/app/model"
	"kapai/app/result"
	"kapai/app/util"
	"strings"
	"time"
)

// Login 登陆
func (s *Service) Login(param *model.LoginParam) result.Result[*model.UserModel] {
	rs := result.NewResult[*model.UserModel]()
	param.Account = strings.TrimSpace(param.Account)
	param.Password = strings.TrimSpace(param.Password)
	if param.Account == "" {
		return rs.WithMsg(result.ErrAccountEmpty)
	}
	if param.Password == "" {
		return rs.WithMsg(result.ErrPasswordEmpty)
	}
	data := s.userDao.GetUser(param.Account)
	if data == nil {
		// 校验账号是否合法
		if !util.CheckAccount(param.Account) {
			return rs.WithMsg(result.ErrAccountCheck)
		}
		// 账号不存在，则创建
		avatar := util.GetAvatar()
		ok := s.userDao.AddUser(&entity.UserEntity{
			Account:    param.Account,
			Name:       param.Account,
			Avatar:     avatar,
			Password:   param.Password,
			Decks:      make([]int, 0),
			Spells:     make([]int, 0),
			CreateTime: time.Now().UnixMilli(),
		})
		if !ok {
			return rs.WithMsg(result.ErrCreateUser)
		}
		token := getJwtToken(param.Account)
		if token == "" {
			return rs.WithMsg(result.ErrToken)
		}
		return rs.WithData(&model.UserModel{
			Account: param.Account,
			Name:    param.Account,
			Avatar:  avatar,
			Token:   token,
			IsNew:   true,
		})
	}
	// 已经存在，则验证密码
	if data.Password != param.Password {
		return rs.WithMsg(result.ErrUserPassword)
	}

	token := getJwtToken(param.Account)
	if token == "" {
		return rs.WithMsg(result.ErrToken)
	}

	return rs.WithData(&model.UserModel{
		Account: data.Account,
		Name:    data.Name,
		Avatar:  data.Avatar,
		Token:   token,
		Decks:   data.Decks,
		Spells:  data.Spells,
		IsNew:   false,
	})
}
