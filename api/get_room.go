package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

var collection *mongo.Collection

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Panic(err)
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		log.Panic(err)
	}
	collection = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_COL"))
}

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
