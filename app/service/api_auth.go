package service

import (
	"github.com/pusher/pusher-http-go/v5"
	"os"
)

var client *pusher.Client

func init() {
	client = &pusher.Client{
		AppID:   os.Getenv("APP_ID"),
		Key:     os.Getenv("APP_KEY"),
		Secret:  os.Getenv("APP_SECRET"),
		Cluster: os.Getenv("APP_CLUSTER"),
	}
}

// Auth 鉴权
func (s *Service) Auth(param []byte) (response []byte, err error) {
	return client.AuthenticatePrivateChannel(param)
}
