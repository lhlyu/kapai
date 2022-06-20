package api

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type Player struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	RoomId string `json:"roomId"`
}

type Room struct {
	Id string `json:"id"`
	// 1 - 等待; 2 - 游戏中
	Status  int     `json:"status"`
	Player1 *Player `json:"player1"`
	Player2 *Player `json:"player2"`
}

var rooms = sync.Map{}
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
	for {
		t := time.Now().UnixNano() / 1e6
		id := strconv.FormatInt(t, 32)
		if _, ok := rooms.Load(id); ok {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		rooms.Store(id, &Room{
			Id:     id,
			Status: 0,
		})
		return id
	}
}

func toJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
