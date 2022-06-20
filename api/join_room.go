package api

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

// JoinRoom 加入房间
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	player := &Player{}
	if err := parseQuery("r", r, player); err != nil {
		NewResult(nil, 2100, err).Fprintf(w)
		return
	}
	if player.Id == "" {
		NewResult(nil, 2101, errors.New("player id miss")).Fprintf(w)
		return
	}
	if player.RoomId == "" {
		NewResult(nil, 2102, errors.New("room id miss")).Fprintf(w)
		return
	}

	room := &Room{}
	if err := collection.FindOne(context.Background(), bson.M{
		"id": player.RoomId,
	}).Decode(room); err != nil {
		NewResult(nil, 2103, err).Fprintf(w)
		return
	}
	if room.Id == "" {
		NewResult(nil, 2104, errors.New("room does not exist")).Fprintf(w)
		return
	}
	if room.Status == 2 {
		NewResult(nil, 2105, errors.New("room is full")).Fprintf(w)
		return
	}
	if room.Player1 == nil && room.Player2 == nil {
		collection.DeleteOne(context.Background(), bson.M{"id": room.Id})
		NewResult(nil, 2106, errors.New("room does not exist")).Fprintf(w)
		return
	}
	if room.Player1 != nil && room.Player2 != nil {
		NewResult(nil, 2107, errors.New("room is full")).Fprintf(w)
		return
	}
	if room.Player1 == nil {
		room.Player1 = player
	}
	if room.Player2 == nil {
		room.Player2 = player
	}

	room.Status = 2

	collection.UpdateOne(context.Background(), bson.M{"id": room.Id}, bson.M{
		"$set": bson.M{
			"status":  room.Status,
			"player1": room.Player1,
			"player2": room.Player2,
		},
	})

	NewResult("ok", 0, nil).Fprintf(w)
	return
}
