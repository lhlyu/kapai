package entity

const (
	ROOM_STATUS_FREE   = iota + 1 // 空闲
	ROOM_STATUS_GAMING            // 游戏中
)

type RoomEntity struct {
	// 秒级时间戳 转 16进制生成
	Id string `json:"id" bson:"id"`
	// 房间状态: 1 - 空闲; 2 - 游戏中
	Status int `json:"status" bson:"status"`
	// 黑方
	Black string `json:"black" bson:"black"`
	// 红方
	Red string `json:"red" bson:"red"`
}
