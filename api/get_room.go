package api

import (
	"errors"
	"net/http"
)

// GetRoom 获取房间信息
func GetRoom(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	val, ok := rooms.Load(id)
	if !ok {
		NewResult(nil, 2300, errors.New("room does not exist")).Fprintf(w)
		return
	}
	NewResult(val, 0, nil).Fprintf(w)
	return
}
