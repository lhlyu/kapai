package api

import (
	"github.com/pusher/pusher-http-go/v5"
	"io"
	"io/ioutil"
	"net/http"
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

func Auth(w http.ResponseWriter, r *http.Request) {
	if !cors(w, r) {
		return
	}
	param, _ := ioutil.ReadAll(r.Body)
	resp, _ := client.AuthenticatePrivateChannel(param)
	io.WriteString(w, string(resp))
}
