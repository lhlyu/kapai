package api

import (
	"net/http"
	"sort"
)

// GetRooms 获取房间
func GetRooms(w http.ResponseWriter, r *http.Request) {

	rs := make([]*Room, 0)

	rooms.Range(func(key, value any) bool {
		if value == nil {
			return true
		}
		rs = append(rs, value.(*Room))
		return true
	})

	sort.SliceStable(rs, func(i, j int) bool {
		return rs[i].Id > rs[j].Id
	})

	NewResult(rs, 0, nil).Fprintf(w)
	return
}
