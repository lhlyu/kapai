package entity

type UserEntity struct {
	// 账号
	Account string `json:"account" bson:"account"`
	// 昵称
	Name string `json:"name" bson:"name"`
	// 头像
	Avatar string `json:"avatar" bson:"avatar"`
	// 密码
	Password string `json:"password" bson:"password"`

	// 卡组
	Decks []int `json:"decks" bson:"decks"`
	// 玩家技能
	Spells []int `json:"spells" bson:"spells"`

	// 创建时间
	CreateTime int64 `json:"createTime" bson:"createTime"`
	// 更新时间
	UpdateTime int64 `json:"updateTime" bson:"updateTime"`

	// 当前所在房间
	RoomId string `json:"roomId" bson:"roomId"`
}
