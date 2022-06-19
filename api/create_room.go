package api

import (
	"encoding/json"
	"errors"
	"github.com/itchyny/base58-go"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Player struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	RoomId   string `json:"roomId"`
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

var encoding = base58.FlickrEncoding

// 解析参数
func parseQuery(field string, r *http.Request, v interface{}) error {
	val := r.URL.Query().Get(field)
	b, err := encoding.Decode([]byte(val))
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
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

	rooms.Store(id, &Room{
		Id:      id,
		Status:  1,
		Player1: player,
	})

	NewResult(player, 0, nil).Fprintf(w)

}
