package api

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
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

	room := &Room{}
	if err := collection.FindOne(context.Background(), bson.M{
		"id": player.RoomId,
	}).Decode(room); err != nil {
		NewResult(nil, 2103, err).Fprintf(w)
		return
	}
	if room.Id == "" {
		NewResult(nil, 0, nil).Fprintf(w)
		return
	}

	if room.Player1.Id == player.Id {
		room.Player1 = nil
	}
	if room.Player2.Id == player.Id {
		room.Player2 = nil
	}

	if room.Player1 == nil && room.Player2 == nil {
		collection.DeleteOne(context.Background(), bson.M{"id": room.Id})
		NewResult("ok", 0, nil).Fprintf(w)
		return
	}

	room.Status = 1

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
