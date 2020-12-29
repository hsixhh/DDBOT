package bilibili

import (
	"errors"
	"github.com/Sora233/Sora233-MiraiGo/utils"
	"github.com/asmcos/requests"
)

const (
	PathRoomInit = "/room/v1/Room/room_init"
)

type RoomInitRequest struct {
	Id int64 `json:"id"`
}

func RoomInit(roomId int64) (*RoomInitResponse, error) {
	url := BPath(PathRoomInit)
	params, err := utils.ToParams(&RoomInitRequest{
		Id: roomId,
	})
	if err != nil {
		return nil, err
	}
	resp, err := requests.Get(url, params)
	if err != nil {
		return nil, err
	}
	rir := new(RoomInitResponse)
	err = resp.Json(rir)
	if err != nil {
		return nil, err
	}
	if rir.GetCode() != 0 {
		return nil, errors.New(rir.Message)
	}
	return rir, nil
}