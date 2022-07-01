package api

import (
	"encoding/json"
	"io"
	"kapai/app/model"
	"kapai/app/result"
	"kapai/app/service"
	"kapai/app/util"
	"net/http"
	"net/url"
)

// Index 入口
func Index(w http.ResponseWriter, r *http.Request) {
	if !cors(w, r) {
		return
	}
	account, ok := auth(w, r)
	if !ok {
		w.WriteHeader(401)
		return
	}

	svc := service.NewService(r.Context(), account)

	io.WriteString(w, route(svc, r))
}

func cors(w http.ResponseWriter, r *http.Request) bool {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Token, Uri")
	}
	if r.Method == "OPTIONS" {
		return false
	}
	return true
}

func auth(w http.ResponseWriter, r *http.Request) (string, bool) {
	uri := r.Header.Get("uri")
	if uri == "login" {
		return "", true
	}
	token := r.Header.Get("token")
	account := service.ParseJwtToken(token)
	if account == "" {
		return "", false
	}
	return account, true
}

func route(svc *service.Service, r *http.Request) string {
	uri := r.Header.Get("uri")

	switch uri {
	case "login":
		param := &model.LoginParam{}
		parseQuery(r, param)
		return svc.Login(param).String()
	case "create.room":
		return svc.CreateRoom().String()
	case "join.room":
		param := &model.RoomParam{}
		parseQuery(r, param)
		return svc.JoinRoom(param).String()
	case "leave.room":
		param := &model.RoomParam{}
		parseQuery(r, param)
		return svc.LeaveRoom(param).String()
	case "get.room":
		roomId := r.URL.Query().Get("r")
		return svc.GetRoom(roomId).String()
	case "get.rooms":
		return svc.GetRooms().String()
	case "update.user":
		param := &model.UpdateUserParam{}
		parseQuery(r, param)
		return svc.UpdateUser(param).String()
	}
	return result.NewResult().WithMsg("未知请求").String()
}

// 解析参数
func parseQuery(r *http.Request, v interface{}) error {
	val := r.URL.Query().Get("k")
	s, err := url.PathUnescape(val)
	if err != nil {
		util.Log.Error().Err(err).Str("s", s).Msg("parseQuery")
		return err
	}
	return json.Unmarshal([]byte(s), v)
}
