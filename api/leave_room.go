package api

import (
	"errors"
	"net/http"
)

// LeaveRoom 离开房间
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	player := &Player{}
	if err := parseQuery("r", r, player); err != nil {
		NewResult(nil, 2200, err).Fprintf(w)
		return
	}
	if player.Id == "" {
		NewResult(nil, 2201, errors.New("player id miss")).Fprintf(w)
		return
	}
	if player.RoomId == "" {
		NewResult(nil, 2202, errors.New("room id miss")).Fprintf(w)
		return
	}
	val, ok := rooms.Load(player.RoomId)
	if !ok {
		NewResult(nil, 2203, errors.New("room does not exist")).Fprintf(w)
		return
	}
	room := val.(*Room)

	if room.Player1.Id == player.Id {
		room.Player1 = nil
	}
	if room.Player2.Id == player.Id {
		room.Player2 = nil
	}

	if room.Player1 == nil && room.Player2 == nil {
		rooms.Delete(room.Id)
		NewResult("ok", 0, nil).Fprintf(w)
		return
	}

	room.Status = 1
	rooms.Store(room.Id, room)
	NewResult("ok", 0, nil).Fprintf(w)
	return
}
