package model

type UserModel struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Token   string `json:"token,omitempty"`

	// 卡组
	Decks []int `json:"decks,omitempty"`
	// 玩家技能
	Spells []int `json:"spells,omitempty"`

	// 是否是新建账号
	IsNew bool `json:"isNew,omitempty"`
}

type Room struct {
	// 秒级时间戳 转 16进制生成
	Id string `json:"id"`
	// 房间状态: 1 - 空闲; 2 - 游戏中
	Status int `json:"status"`
	// 黑方
	Black *UserModel `json:"black,omitempty"`
	// 红方
	Red *UserModel `json:"red,omitempty"`
}
