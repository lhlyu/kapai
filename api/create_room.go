package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type Player struct {
	Id     string `json:"id" bson:"id"`
	Name   string `json:"name" bson:"name"`
	RoomId string `json:"roomId" bson:"roomId"`
}

type Room struct {
	Id string `json:"id" bson:"id"`
	// 1 - 等待; 2 - 游戏中
	Status  int     `json:"status" bson:"status"`
	Player1 *Player `json:"player1" bson:"player1"`
	Player2 *Player `json:"player2" bson:"player2"`
}

var lock = sync.Mutex{}

// 解析参数
func parseQuery(field string, r *http.Request, v interface{}) error {
	val := r.URL.Query().Get(field)
	s, err := url.PathUnescape(val)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(s), v)
}

// 生成房间id
func getRoomId() string {
	t := time.Now().UnixNano() / 1e6
	id := strconv.FormatInt(t, 32)
	return id
}

// CreateRoom 创建房间
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	player := &Player{}
	if err := parseQuery("r", r, player); err != nil {
		NewResult(nil, 2000, err).Fprintf(w)
		return
	}
	if player.Id == "" {
		NewResult(nil, 2001, errors.New("player id miss")).Fprintf(w)
		return
	}
	// 获取房间id
	id := getRoomId()
	player.RoomId = id
	_, err := collection.InsertOne(context.Background(), &Room{
		Id:      id,
		Status:  1,
		Player1: player,
		Player2: nil,
	})
	if err != nil {
		NewResult(nil, 2002, err).Fprintf(w)
		return
	}
	NewResult(player, 0, nil).Fprintf(w)
}
