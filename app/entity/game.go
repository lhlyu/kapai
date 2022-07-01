package entity

const (
	GAME_KIND_DECK  = iota + 1 // 卡牌 [10000, )
	GAME_KIND_SPELL            // 技能 [1, 10000)
)

type GameEntity struct {
	Id   int `json:"id" bson:"id"`
	kind int `json:"kind" bson:"kind"`
	// 名字
	Name string `json:"name" bson:"name"`
	// 图片
	Image string `json:"image" bson:"image"`
	// 规格
	Spec string `json:"spec" bson:"spec"`
}
