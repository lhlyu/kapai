package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// GetRoom 获取房间信息
func GetRoom(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	room := &Room{}
	err := collection.FindOne(context.Background(), bson.M{
		"id": id,
	}).Decode(room)
	if err != nil {
		NewResult(nil, 2300, err).Fprintf(w)
		return
	}
	if room.Id == "" {
		NewResult(nil, 0, nil).Fprintf(w)
		return
	}
	NewResult(room, 0, nil).Fprintf(w)
	return
}
