package api

import (
	"errors"
	"log"
	"net/http"
)

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

	v, _ := rooms.Load(id)

	log.Println("【创建房间】:", toJson(v))
	NewResult(player, 0, nil).Fprintf(w)

}
