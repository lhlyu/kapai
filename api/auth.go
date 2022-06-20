package api

import (
	"encoding/json"
	"fmt"
	"github.com/pusher/pusher-http-go/v5"
	"io/ioutil"
	"log"
	"net/http"
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

func Auth(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if r.Method == "OPTIONS" {
		return
	}
	// 响应http code
	w.WriteHeader(200)

	params, _ := ioutil.ReadAll(r.Body)
	log.Println("params:", string(params))
	response, err := Client.AuthenticatePrivateChannel(params)
	fmt.Println("response:", string(response))
	NewResult(string(response), 1000, err).Fprintf(w)
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResult(data interface{}, code int, err error) *Result {
	if err != nil {
		return &Result{
			Code: code,
			Msg:  err.Error(),
		}
	}
	return &Result{
		Data: data,
	}
}

func (r *Result) Fprintf(w http.ResponseWriter) {
	fmt.Fprintf(w, r.ToJson())
}

func (r *Result) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}
