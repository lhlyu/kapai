package pkg

import (
	"github.com/pusher/pusher-http-go/v5"
	"os"
)

var Client *pusher.Client

func init() {
	Client = &pusher.Client{
		AppID:   os.Getenv("APP_ID"),
		Key:     os.Getenv("APP_KEY"),
		Secret:  os.Getenv("APP_SECRET"),
		Cluster: os.Getenv("APP_CLUSTER"),
	}
}
