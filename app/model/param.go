package model

// LoginParam 登陆参数
type LoginParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// UpdateUserParam 更新用户参数
type UpdateUserParam struct {
	// 昵称
	Name *string `json:"name,omitempty"`
	// 头像
	Avatar *string `json:"avatar,omitempty"`
	// 密码
	Password *string `json:"password,omitempty"`
	// 卡组
	Decks *[]int `json:"decks,omitempty"`
	// 玩家技能
	Spells *[]int `json:"spells,omitempty"`
}

type RoomParam struct {
	// 房间iD
	RoomId string `json:"room_id"`
}
