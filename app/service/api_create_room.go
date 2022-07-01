package service

import (
	"kapai/app/entity"
	"kapai/app/result"
	"strconv"
	"time"
)

func (s *Service) CreateRoom() result.Result {
	rs := result.NewResult()

	var roomId string
	for i := 0; i < 10; i++ {
		roomId = getRoomId()
		if s.roomDao.GetRoom(roomId) == nil {
			time.Sleep(time.Second)
			continue
		}
	}
	if roomId == "" {
		return rs.WithMsg("创建房间失败")
	}

	ok := s.roomDao.AddRoom(&entity.RoomEntity{
		Id:     roomId,
		Status: 1,
	})
	if !ok {
		return rs.WithMsg("创建房间异常")
	}

	return rs.WithData(roomId)
}

func getRoomId() string {
	return strconv.FormatInt(time.Now().Unix(), 16)
}
