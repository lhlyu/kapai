package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

// GetRooms 获取房间
func GetRooms(w http.ResponseWriter, r *http.Request) {
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		NewResult(nil, 2400, err).Fprintf(w)
		return
	}

	rs := make([]*Room, 0)
	for cur.Next(context.Background()) {
		room := &Room{}
		cur.Decode(room)
		if room.Id != "" {
			rs = append(rs, room)
		}
	}
	log.Println("【当前房间数】:", len(rs))
	NewResult(rs, 0, nil).Fprintf(w)
	return
}
